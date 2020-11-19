package help

import (
	"io/ioutil"
	"os"
)

// ReadFile : Reads a file from a given path.
func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	byteFile, _ := ioutil.ReadAll(f)
	return string(byteFile), nil
}
