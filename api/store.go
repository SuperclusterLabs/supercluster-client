package api

import (
	"errors"
	"log"
	"os"
	"sort"
	"time"
)

const StoreName = "store"

var ErrNotFound = errors.New("File does not exist")
var ErrFileExists = errors.New("File already exists")
var ErrRequestUnmarshalled = errors.New("Request could not be unmarshalled")
var ErrCannotCreate = errors.New("File could not be created")

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
		log.Println("Could not create file: ", err.Error())
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

	_, err = f.WriteString(contents)
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

func (s *store) Modify(name, contents string) (*file, error) {
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

	var modified *file = s.files[name]
	modified.Contents = contents

	return modified, nil
}

func (s *store) Delete(name string) error {
	err := os.Remove(StoreName + "/" + name)
	if err != nil {
		log.Println("Could not delete file: ", err.Error())
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
