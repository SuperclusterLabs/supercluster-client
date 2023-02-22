package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/runtime"

	"github.com/gin-gonic/gin"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

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
	var data map[string]string

	// This is a hack to track metadata for a file. Since a dir is a file
	// containing file info, we can use it to track file metadata.
	// N.B: IPFS only stores name, size (bytes), and cid
	resp, err := http.Post("http://localhost:9095/api/v0/add?wrap-with-directory=true", ctx.ContentType(), ctx.Request.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(string(body)), &data); err != nil {
		return nil, errors.New("Could not add file to ipfs-cluster")
	}

	// TODO: figure out a way to embed created time/creator info
	// into ipfs file description
	new := &model.File{
		Cid:  data["cid"],
		Name: name,
		Size: int64(len(contents)),
		// TODO: is pin type only one of 2 options?
		PinType: "recursive",
	}

	return new, nil
}

func (s *IPFSClusterStore) Modify(ctx *gin.Context, name, contents string) (*model.File, error) {
	return nil, nil
}

	p := path.New(cid)
	err := s.ipfsApi.Pin().Rm(ctx, p)
	if err != nil {
		log.Println("Could not remove file ", err.Error())
		return err
	}
func (s *IPFSClusterStore) Delete(ctx *gin.Context, cid string) error {

	return nil
}

func (s *IPFSClusterStore) DeleteAll(ctx *gin.Context) error {
	fs, err := s.List(ctx)
	if err != nil {
		log.Println("Could not fetch pinned files ", err.Error())
		return err
	}

	for _, f := range fs {
		if f.PinType == "recursive" {
			err := s.ipfsApi.Pin().Rm(ctx, path.New(f.Cid))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *IPFSClusterStore) List(ctx *gin.Context) ([]model.File, error) {
	files := make([]model.File, 0)

	pins, err := s.ipfsApi.Pin().Ls(ctx)
	if err != nil {
		return nil, err
	}

	// since all files are directories, grab name from them
	for p := range pins {
		dir := false
		es, err := s.ipfsApi.Unixfs().Ls(ctx, p.Path())
		if err != nil {
			return nil, err
		}

		// ignore indirect pins as they are necessarily files for now

		for e := range es {
			dir = true
			files = append(files, model.File{
				Cid:     e.Cid.String(),
				Name:    e.Name,
				Size:    int64(e.Size),
				PinType: "indirect",
			})
		}

		if dir {
			files = append(files, model.File{
				Cid:     p.Path().Cid().String(),
				PinType: p.Type(),
			})
		}
	}

	return files, nil
}

	resp, err := http.Post("http://localhost:5001/api/v0/id", "application/json", nil)
func (s *IPFSClusterStore) GetInfo(ctx *gin.Context) (*P2PNodeInfo, error) {
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

	err := s.ipfsApi.Pin().Add(ctx, path.New(c))
func (s *IPFSClusterStore) PinFile(ctx *gin.Context, cid string) error {
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
