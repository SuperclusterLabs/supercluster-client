package store

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	model "github.com/SuperclusterLabs/supercluster-client/model"
	util "github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/gin-gonic/gin"
)

const StoreName = "store"

type osStore struct {
	files map[string]*model.File
}

func NewOSStore() (P2PStore, error) {
	s := &osStore{
		files: make(map[string]*model.File),
	}
	err := os.Mkdir(StoreName, 0777)
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}
	return s, nil
}

func (s *osStore) Create(_ *gin.Context, name string, contents []byte) (*model.File, error) {
	new := &model.File{
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
	if _, ok := s.files[name]; ok {
		log.Println("Could not create file: ", util.ErrFileExists.Error())
		return nil, util.ErrFileExists
	}
	s.files[name] = new

	f, err := os.Create(StoreName + "/" + name)
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	_, err = f.WriteString(string(contents))
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}

	f.Close()
	if err != nil {
		log.Println("Could not close file: ", err.Error())
		return nil, err
	}

	return new, nil
}

func (s *osStore) Modify(ctx *gin.Context, name, contents string) (*model.File, error) {
	filename := StoreName + "/" + name
	if err := os.Truncate(filename, 0); err != nil {
		log.Println("Failed to truncate: ", err)
		return nil, err
	}

	f, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		log.Println("Could not modify file: ", err.Error())
		return nil, err
	}

	_, err = f.WriteString(contents)
	if err != nil {
		log.Println("Could not modify file: ", err.Error())
		return nil, err
	}

	return nil, nil
}

func (s *osStore) Delete(ctx *gin.Context, name string) error {
	err := os.Remove(StoreName + "/" + name)
	if err != nil {
		log.Println("Could not delete file: ", err.Error())
		return err
	}
	return nil
}

func (s *osStore) DeleteAll(ctx *gin.Context) error {
	fs, err := s.List(ctx)
	if err != nil {
		log.Println("Could not fetch pinned files ", err.Error())
		return err
	}

	for _, f := range fs {
		err := os.Remove(StoreName + "/" + f.Name)
		if err != nil {
			log.Println("Could not delete file: ", err.Error())
			return err
		}
	}

	return nil
}

func (s *osStore) List(ctx *gin.Context) ([]model.File, error) {
	files := make([]model.File, 0)
	existing, err := ioutil.ReadDir(StoreName)
	if err != nil {
		log.Fatal("`store` dir could not be accessed:", err.Error())
		return nil, err
	}

	for _, fInfo := range existing {
		f := model.File{
			Name:      fInfo.Name(),
			CreatedAt: fInfo.ModTime().Unix(),
		}

		if _, ok := s.files[f.Name]; !ok {
			s.files[f.Name] = &f
		}
	}

	for _, f := range s.files {
		files = append(files, *f)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].CreatedAt < files[j].CreatedAt
	})

	return files, nil
}

func (s *osStore) GetInfo(ctx *gin.Context) (*P2PNodeInfo, error) {
	return nil, errors.New("Not supported")
}

func (s *osStore) PinFile(ctx *gin.Context, c string) error {
	return errors.New("Not supported")
}

func (s *osStore) ConnectPeer(ctx *gin.Context, addr ...string) error {
	return errors.New("Not supported")
}
