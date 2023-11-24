package file

import (
	"os"
	"path/filepath"

	"go.k6.io/k6/js/modules"
)

type FILE struct{}

func init() {
	modules.Register("k6/x/file", new(FILE))
}

func (*FILE) CreateDir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (*FILE) CreateFile(path string, filename string) error {
	f, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
