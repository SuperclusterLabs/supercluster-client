package store

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SuperclusterLabs/supercluster-client/model"

	"github.com/gin-gonic/gin"
	ipfsFiles "github.com/ipfs/go-ipfs-files"
	shell "github.com/ipfs/go-ipfs-http-client"
	path "github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// TODO: we should transition this to (what?) metadata
type IPFSStore struct {
	// abstraction that maps file names to file structs
	// reqd for file modifications, etc
	files   map[string]*model.File
	ipfsApi *shell.HttpApi
}

var _ P2PStore = (*IPFSStore)(nil)

func NewIPFSStore() (*IPFSStore, error) {
	api, err := shell.NewURLApiWithClient("localhost:5001", &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	})
	if err != nil {
		return nil, err
	}
	s := &IPFSStore{
		files:   make(map[string]*model.File),
		ipfsApi: api,
	}

	return s, nil
}

func (s *IPFSStore) Get(ctx *gin.Context, cid string) ([]byte, *model.File, error) {
	return nil, nil, nil
}

func (s *IPFSStore) Create(ctx *gin.Context, name string, contents []byte) (*model.File, error) {
	// This is a hack to track metadata for a file. Since a dir is a file
	// containing file info, we can use it to track file metadata.
	// N.B: IPFS only stores name, size (bytes), and cid
	f := ipfsFiles.NewBytesFile(contents)
	cid, err := s.ipfsApi.Unixfs().
		Add(ctx, ipfsFiles.NewMapDirectory(map[string]ipfsFiles.Node{
			name: f,
		}))
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	err = s.ipfsApi.Pin().Add(ctx, cid)
	if err != nil {
		log.Println("Could not pin file: ", err.Error())
		return nil, err
	}

	// TODO: figure out a way to embed created time/creator info
	// into ipfs file description
	new := &model.File{
		Cid:  cid.Cid().String(),
		Name: name,
		Size: int64(len(contents)),
		// TODO: is pin type only one of 2 options?
		PinType: "recursive",
	}

	return new, nil
}

func (s *IPFSStore) Modify(ctx *gin.Context, name, contents string) (*model.File, error) {
	return nil, nil
}

func (s *IPFSStore) Delete(ctx *gin.Context, cid string) error {
	p := path.New(cid)
	err := s.ipfsApi.Pin().Rm(ctx, p)
	if err != nil {
		log.Println("Could not remove file ", err.Error())
		return err
	}

	return nil
}

func (s *IPFSStore) DeleteAll(ctx *gin.Context) error {
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

func (s *IPFSStore) List(ctx *gin.Context) ([]model.File, error) {
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

func (s *IPFSStore) GetInfo(ctx *gin.Context) (*P2PNodeInfo, error) {
	resp, err := http.Post("http://localhost:5001/api/v0/id", "application/json", nil)
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

func (s *IPFSStore) PinFile(ctx *gin.Context, c string) error {
	err := s.ipfsApi.Pin().Add(ctx, path.New(c))
	return err
}

func (s *IPFSStore) ConnectPeer(ctx *gin.Context, addrs ...string) error {
	var ms []ma.Multiaddr
	for _, a := range addrs {
		m, err := ma.NewMultiaddr(a)
		if err != nil {
			return err
		}

		ms = append(ms, m)
	}
	err := s.ipfsApi.Swarm().Connect(ctx, peer.AddrInfo{
		Addrs: ms,
	})
	return err
}
