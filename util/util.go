package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var DirName string = ".supercluster"

func GetConfDir() string {
	hDir, err := os.UserHomeDir()
	if err != nil {
		// this shouldn't fire on any OS' we care about
		panic(err)
	}
	return hDir + "/" + DirName
}

func ReadJSONFile(filename string) (map[string]interface{}, error) {
	// Read the file contents into a byte slice
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a map[string]string
	var data map[string]interface{}
	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
