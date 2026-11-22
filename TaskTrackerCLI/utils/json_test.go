package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateOrGetFile(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("resultado '%s', esperado nil", err)
	}

	path := filepath.Join(home, "tasks.json")
	file, err := CreateOrGetFile(path)

	if err != nil {
		t.Errorf("resultado '%s', esperado nil", err)
	}
	if file == nil {
		t.Errorf("resultado '%s', esperado nome", file.Name())
	}
}
