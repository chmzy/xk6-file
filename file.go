package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.k6.io/k6/js/modules"
)

type FILE struct{}

func init() {
	modules.Register("k6/x/file", new(FILE))
}

func (*FILE) CreateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("CreateDir: path is empty")
	}

	var validatedPath string = path

	if strings.HasPrefix(path, "~/") {
		validatedPath = filepath.Join(os.Getenv("HOME"), path[1:])
	}

	if strings.HasPrefix(path, "$HOME/") {
		validatedPath = os.ExpandEnv(path)
	}

	if strings.HasPrefix(path, "./") {
		wd, _ := os.Getwd()
		validatedPath = filepath.Join(wd, path[1:])
	}

	if !filepath.IsAbs(validatedPath) {
		return "", fmt.Errorf("CreateDir: define absolute or relative path to dir")
	}

	if err := os.MkdirAll(validatedPath, os.ModePerm); err != nil {
		return "", err
	}
	return validatedPath, nil
}
