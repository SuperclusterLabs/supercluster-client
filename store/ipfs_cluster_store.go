package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/runtime"

	"github.com/gin-gonic/gin"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

type objects struct {
	ObjectsList []struct {
		Hash  string `json:"Hash"`
		Links []struct {
			Name   string `json:"Name"`
			Hash   string `json:"Hash"`
			Size   int    `json:"Size"`
			Type   int    `json:"Type"`
			Target string `json:"Target"`
		} `json:"Links"`
	} `json:"Objects"`
}

type IPFSClusterStore struct {
	clusterSvcPort string

	*IPFSStore
}

var _ P2PStore = (*IPFSClusterStore)(nil)

func NewIPFSClusterStore() (*IPFSClusterStore, error) {
	is, err := NewIPFSStore()
	if err != nil {
		return nil, err
	}
	s := &IPFSClusterStore{
		IPFSStore: is,
	}

	return s, nil
}

func (s *IPFSClusterStore) Create(ctx *gin.Context, name string, contents []byte) (*model.File, error) {

	// This is a hack to track metadata for a file. Since a dir is a file
	// containing file info, we can use it to track file metadata.
	// N.B: IPFS only stores name, size (bytes), and cid
	createEp := "/add?wrap-with-directory=true"
	clsResp, err := makeClusterSvcRequest(ctx, createEp)
	if err != nil {
		return nil, err
	}

	if cid, ok := clsResp["Hash"]; ok {
		// TODO: figure out a way to embed created time/creator info
		// into ipfs file description
		new := &model.File{
			Cid:  cid.(string),
			Name: name,
			Size: int64(len(contents)),
			// TODO: is pin type only one of 2 options?
			PinType: "recursive",
		}

		return new, nil
	}

	// FIXME: this is weird
	errStr, err := json.Marshal(clsResp)
	if err != nil {
		return nil, err
	}
	return nil, errors.New(string(errStr))
}

func (s *IPFSClusterStore) Modify(ctx *gin.Context, name, contents string) (*model.File, error) {
	return nil, nil
}

func (s *IPFSClusterStore) Delete(ctx *gin.Context, cid string) error {
	createEp := "/pin/rm/" + cid
	_, err := makeClusterSvcRequest(ctx, createEp)

	return err
}

func (s *IPFSClusterStore) DeleteAll(ctx *gin.Context) error {
	fs, err := s.List(ctx)
	if err != nil {
		log.Println("Could not fetch pinned files ", err.Error())
		return err
	}

	for _, f := range fs {
		if f.PinType == "recursive" {
			err := s.Delete(ctx, f.Cid)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *IPFSClusterStore) List(ctx *gin.Context) ([]model.File, error) {
	files := make([]model.File, 0)

	createEp := "/pin/ls"
	clsResp, err := makeClusterSvcRequest(ctx, createEp)
	if err != nil {
		return nil, err
	}

	if cids, ok := clsResp["Keys"]; ok {
		cidsMap := cids.(map[string]any)
		for k := range cidsMap {
			// ipfs cluster doesn't pin the contents of directories, rather just the dirs
			// so we need to fetch the contents
			lsEp := "/ls/" + k
			clsResp, err := makeClusterSvcRequest(ctx, lsEp)
			if err != nil {
				return nil, err
			}

			jsonString, _ := json.Marshal(clsResp)
			var o objects
			err = json.Unmarshal([]byte(jsonString), &o)
			if err != nil {
				return nil, err
			}
			if len(o.ObjectsList) == 0 {
				return nil, errors.New("empty object list")
			}
			firstObject := o.ObjectsList[0]
			if len(firstObject.Links) == 0 {
				log.Println("empty links for " + k)
				continue
			}
			l := firstObject.Links[0]

			// TODO: figure out a way to embed created time/creator info
			// into ipfs file description
			f := model.File{
				Cid:  l.Hash,
				Name: l.Name,
				Size: int64(l.Size),
			}

			files = append(files, f)
		}
		return files, nil
	}

	errStr, _ := json.Marshal(clsResp)
	return nil, errors.New(string(errStr))
}

func (s *IPFSClusterStore) GetInfo(ctx *gin.Context) (*P2PNodeInfo, error) {
	// FIXME: make consistent with other API calls
	c := ctx.Param("clusterId")
	u, err := getClusterURL(c)

	resp, err := http.Post(u+"/id", "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ar P2PNodeInfo
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (s *IPFSClusterStore) PinFile(ctx *gin.Context, cid string) error {
	createEp := "/pin/add/" + cid
	_, err := makeClusterSvcRequest(ctx, createEp)

	return err
}

func getClusterURL(clusterId string) (string, error) {
	icp, err := runtime.GlobalRuntime.GetProcess(clusterId)
	if err != nil {
		return "", err
	}
	p, err := icp.GetPort()
	if err != nil {
		return "", err
	}
	return "http://localhost:" + p + "/api/v0", nil
}

// endpoint should start with `/`
func makeClusterSvcRequest(ctx *gin.Context, endpoint string) (map[string]interface{}, error) {
	c := ctx.Param("clusterId")
	u, err := getClusterURL(c)
	if err != nil {
		return nil, err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileWriter, err := writer.CreateFormFile("file", "filename")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u+endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("ipfs service err status code: " + strconv.Itoa(resp.StatusCode))
	}

	var clsResp map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&clsResp)
	if err != nil {
		return nil, err
	}

	return clsResp, nil
}
