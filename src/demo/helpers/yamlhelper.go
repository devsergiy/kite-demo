package helpers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetYamlContent(path string) ([]byte, error) {
	filename, _ := filepath.Abs(path)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("No such file %s", filename)
		return nil, err
	}

	return content, nil
}
