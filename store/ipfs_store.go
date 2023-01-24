package store

import (
	"bytes"
	"context"
	"log"
	"sort"
	"time"

	"github.com/SuperclusterLabs/supercluster-client/model"
	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
)

// TODO: we should transition this to (what?) metadata
type IpfsStore struct {
	// abstraction that maps file names to file structs
	// reqd for file modifications, etc
	files   map[string]*model.File
	ipfsApi *shell.Shell
}

var _ P2PStore = (*IpfsStore)(nil)

func NewIpfsStore() P2PStore {
	s := &IpfsStore{
		files:   make(map[string]*model.File),
		ipfsApi: shell.NewShell("localhost:5001"),
	}

	return s
}

func (s *IpfsStore) Create(ctx *gin.Context, name string, contents []byte) (*model.File, error) {
	// Allow the same file to be pinned again for now, no harm
	// if _, ok := s.files[name]; ok {
	// log.Println("Could not create file: ", ErrFileExists.Error())
	// return nil, ErrFileExists
	// }

	cid, err := s.ipfsApi.Add(bytes.NewReader(contents))
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	// TODO: does adding above automatically pin? Do we only need
	// one of these 2 calls?
	err = s.ipfsApi.Pin(cid)
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	new := &model.File{
		Cid:       cid,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
	s.files[name] = new

	return new, nil
}

func (s *IpfsStore) Modify(ctx context.Context, name, contents string) (*model.File, error) {
	return nil, nil
}

func (s *IpfsStore) Delete(ctx context.Context, cId string) error {
	// f := s.files[name]
	err := s.ipfsApi.Unpin(cId)
	if err != nil {
		log.Println("Could not remove file ", err.Error())
		return err
	}

	// delete(s.files, name)

	return nil
}

func (s *IpfsStore) List(ctx context.Context) ([]model.File, error) {
	files := make([]model.File, 0)

	for _, f := range s.files {
		files = append(files, *f)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].CreatedAt < files[j].CreatedAt
	})

	return files, nil
}

func (s *IpfsStore) GetInfo(ctx context.Context) (*util.AddrsResponse, error) {
	n, err := s.ipfsApi.ID()
	if err != nil {
		return nil, err
	}
	return &util.AddrsResponse{
		ID:    n.ID,
		Addrs: n.Addresses,
	}, nil
}

func (s *IpfsStore) PinFile(ctx *gin.Context, c string) error {
	err := s.ipfsApi.Pin(c)
	return err
}

func (s *IpfsStore) ConnectPeer(ctx *gin.Context, addr ...string) error {
	err := s.ipfsApi.SwarmConnect(ctx, addr...)
	return err
}
