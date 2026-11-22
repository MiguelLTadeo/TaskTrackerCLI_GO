package utils

import (
	"os"
)

func CreateOrGetFile(path string) (*os.File, error) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.Open(path)
}
