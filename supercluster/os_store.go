package supercluster

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

const StoreName = "store"

type osStore struct {
	files map[string]*file
}

func newOSStore() (Store, error) {
	s := &osStore{
		files: make(map[string]*file),
	}
	err := os.Mkdir(StoreName, 0777)
	if err != nil {
		log.Println("Could not create file: ", err.Error())
		return nil, err
	}
	return s, nil
}

func (s *osStore) Create(_ *gin.Context, name string, contents []byte) (*file, error) {
	new := &file{
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
	if _, ok := s.files[name]; ok {
		log.Println("Could not create file: ", ErrFileExists.Error())
		return nil, ErrFileExists
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

func (s *osStore) Modify(ctx context.Context, name, contents string) (*file, error) {
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

func (s *osStore) Delete(ctx context.Context, name string) error {
	err := os.Remove(StoreName + "/" + name)
	if err != nil {
		log.Println("Could not delete file: ", err.Error())
		return err
	}
	delete(s.files, name)
	return nil
}

func (s *osStore) List(ctx context.Context) ([]file, error) {
	files := make([]file, 0)
	existing, err := ioutil.ReadDir(StoreName)
	if err != nil {
		log.Fatal("`store` dir could not be accessed:", err.Error())
		return nil, err
	}

	for _, fInfo := range existing {
		f := file{
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
