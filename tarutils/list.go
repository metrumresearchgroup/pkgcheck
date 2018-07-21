package tarutils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dpastoor/goutils"
	"github.com/r-infra/pkgcheck/rcmd"
	"github.com/spf13/afero"
)

// ListTars lists the path of all tar.gz files recursively from the root dir
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

// CopyPkgTars copies all package tars to a dest directory, creating the dest if it does not exist.
// it returns an array of tarballs copied and error
// heuristically, it determines whether a tarball is a package by assuming any tarball
// with the R package tarball naming convention of <pkgname>_<version>.tar.gz is a package,
// and if not, it inspects the tarball for a DESCRIPTION file.
func CopyPkgTars(fs afero.Fs, root string, dest string, fm rcmd.FilterMap) ([]string, error) {
	var copied []string
	tars := ListTars(root)
	ok, err := afero.DirExists(fs, dest)
	if err != nil {
		return copied, err
	}
	if !ok {
		err = fs.MkdirAll(dest, 0755)
		if err != nil {
			return copied, err
		}
	}
	for _, t := range tars {
		tarball := filepath.Base(t)
		tarball = strings.TrimSuffix(tarball, ".tar.gz")
		// package tarball stored as <package>_<version>.tar.gz
		var pm rcmd.Package
		packageVersion := strings.SplitN(tarball, "_", 2)
		if len(packageVersion) < 2 {
			// if not a package this will come back as empty for name/version
			pm = PackageInfo(t)
		} else {
			pm.Name = packageVersion[0]
			pm.Version = packageVersion[1]
		}
		ok := rcmd.ShouldCheck(pm.Name, fm)
		if ok {
			copyPath := filepath.Clean(filepath.Join(dest, fmt.Sprintf("%s_%s.tar.gz", pm.Name, pm.Version)))
			_, err := goutils.CopyFS(fs, t, copyPath)
			if err != nil {
				// TODO: hmm don't know what the best thing to do logging wise
			}
			copied = append(copied, copyPath)
		}
	}

	return copied, nil
}
