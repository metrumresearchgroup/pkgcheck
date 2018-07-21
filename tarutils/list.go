package tarutils

import (
	"os"
	"path/filepath"
	"strings"
)

// ListTars lists the absolute path of all tar.gz files recursively from the root dir
func ListTars(root string) []string {
	tars := []string{}
	visit := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".tar.gz") {
			tars = append(tars, path)
		}
		return nil
	}
	filepath.Walk(root, visit)
	return tars
}
