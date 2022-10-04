package api

import (
	"errors"
	"log"
	"os"
	"sort"
	"time"
)

const StoreName = "store"

var ErrNotFound = errors.New("not found")

type Store interface {
	Create(name, contents string) (*file, error)
	Modify(name, contents string) (*file, error)
	Delete(name string) error
	List() []file
}

type store struct {
	files map[string]*file
}

type file struct {
	// Name of the file
	Name string `json:"name"`
	// The contents of the file
	Contents string `json:"contents"`
	// Unix timestamp of creation
	CreatedAt int64 `json:"createdAt"`
}

func newStore() (Store, error) {
	s := &store{
		files: make(map[string]*file),
	}
	err := os.Mkdir(StoreName, 0777)
	if err != nil {
		log.Fatal("Could not create file: ", err.Error())
		return nil, err
	}
	return s, nil
}

func (s *store) Create(name, contents string) (*file, error) {
	new := &file{
		Name:      name,
		Contents:  contents,
		CreatedAt: time.Now().Unix(),
	}
	s.files[name] = new

	f, err := os.Create(StoreName + "/" + name)
	if err != nil {
		log.Fatal("Could not create file: ", err.Error())
		return nil, err
	}

	_, err = f.WriteString(contents)
	if err != nil {
		log.Fatal("Could not create file: ", ErrNotFound.Error())
		return nil, err
	}

	f.Close()
	if err != nil {
		log.Fatal("Could not close file: ", ErrNotFound.Error())
		return nil, err
	}

	return new, nil
}

func (s *store) Modify(name, contents string) (*file, error) {
	filename := StoreName + "/" + name
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
		return nil, err
	}

	f, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("Could not modify file: ", ErrNotFound.Error())
		return nil, ErrNotFound
	}

	_, err = f.WriteString(contents)
	if err != nil {
		log.Fatal("Could not modify file: ", ErrNotFound.Error())
		return nil, err
	}

	var modified *file = s.files[name]
	modified.Contents = contents

	return modified, nil
}

func (s *store) Delete(name string) error {
	err := os.Remove(StoreName + "/" + name)
	if err != nil {
		log.Fatal("Could not modify file: ", ErrNotFound.Error())
		return err
	}
	delete(s.files, name)
	return nil
}

func (s *store) List() []file {
	files := make([]file, 0)

	for _, f := range s.files {
		files = append(files, *f)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].CreatedAt < files[j].CreatedAt
	})

	return files
}
