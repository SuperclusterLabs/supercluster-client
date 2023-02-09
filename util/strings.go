package util

import "os"

var DirName string = ".supercluster"

func GetConfDir() string {
	hDir, err := os.UserHomeDir()
	if err != nil {
		// this shouldn't fire on any OS' we care about
		panic(err)
	}
	return hDir + "/" + DirName
}
