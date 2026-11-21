package utils

import (
	"fmt"
	"os"
)

func CreateOrGetFile(path string) (*os.File, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		fmt.Printf("file: %v\n", file)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Open(path)
}
