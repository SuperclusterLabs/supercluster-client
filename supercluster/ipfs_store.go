package supercluster

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	files "github.com/ipfs/go-ipfs-files"
	icorepath "github.com/ipfs/interface-go-ipfs-core/path"
)

// TODO: we should transition this to (what?) metadata
type ipfsStore struct {
	// abstraction that maps file names to file structs
	// reqd for file modifications, etc
	files map[string]*file
}

var _ Store = (*ipfsStore)(nil)

func newIpfsStore() (*ipfsStore, error) {
	s := &ipfsStore{
		files: make(map[string]*file),
	}

	return s, nil
}

func (s *ipfsStore) Create(ctx *gin.Context, name string, contents []byte) (*file, error) {
	ipfs := *getCoreAPIInstance()

	// Allow the same file to be pinned again for now, no harm
	// if _, ok := s.files[name]; ok {
	// log.Println("Could not create file: ", ErrFileExists.Error())
	// return nil, ErrFileExists
	// }

	peerCid, err := ipfs.Unixfs().Add(ctx, files.NewBytesFile(contents))
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	new := &file{
		ID:        peerCid.Cid().String(),
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
	s.files[name] = new

	// TODO: does adding above automatically pin? Do we only need
	// one of these 2 calls?
	ipfs.Pin().Add(ctx, peerCid)

	return new, nil
}

func (s *ipfsStore) Modify(ctx context.Context, name, contents string) (*file, error) {
	// if _, ok := s.files[name]; !ok {
	// log.Println("Could not modify file: ", ErrNotFound.Error())
	// return nil, ErrNotFound
	// }

	// ipfs := *getCoreAPIInstance()

	// remove old cid
	// f := s.files[name]
	// icp := icorepath.New(f.ID)
	// err := ipfs.Pin().Rm(ctx, icp)
	// if err != nil {
	// log.Println("Could not remove old cid ", err.Error())
	// return nil, err
	// }

	// upload+pin file and update cid
	// peerCid, err := ipfs.Unixfs().Add(ctx, files.NewBytesFile([]byte(contents)))
	// if err != nil {
	// log.Println("Could not create file: ", err.Error())
	// return nil, err
	// }

	// f.ID = peerCid.Cid().String()
	// f.Contents = contents
	// ipfs.Pin().Add(ctx, peerCid)

	return nil, nil
}

func (s *ipfsStore) Delete(ctx context.Context, cId string) error {
	// if _, ok := s.files[cId]; !ok {
	// log.Println("Could not modify file: ", ErrNotFound.Error())
	// return ErrNotFound
	// }

	ipfs := *getCoreAPIInstance()

	// f := s.files[name]
	icp := icorepath.New(cId)
	err := ipfs.Pin().Rm(ctx, icp)
	if err != nil {
		log.Println("Could not remove file ", err.Error())
		return err
	}

	// delete(s.files, name)

	return nil
}

func (s *ipfsStore) List(ctx context.Context) ([]file, error) {
	files := make([]file, 0)

	for _, f := range s.files {
		files = append(files, *f)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].CreatedAt < files[j].CreatedAt
	})

	return files, nil
}
