package tarutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver"

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
		var pm PackageTarball
		packageVersion := strings.SplitN(tarball, "_", 2)
		if len(packageVersion) < 2 {
			// if not a package this will come back as empty for name/version
			pm.Details = PackageInfo(t)
		} else {
			pm.Details.Name = packageVersion[0]
			pm.Details.Version = packageVersion[1]
			pm.PackratHashed = true
		}
		ok := rcmd.ShouldCheck(pm.Details.Name, fm)
		if ok {
			copyPath := filepath.Clean(filepath.Join(dest, fmt.Sprintf("%s_%s.tar.gz", pm.Details.Name, pm.Details.Version)))
			if !pm.PackratHashed {
				_, err := goutils.CopyFS(fs, t, copyPath)
				if err != nil {
					// TODO: hmm don't know what the best thing to do logging wise
				}
			} else {
				ok, err := CopyPackratTarball(pm.Details, t, dest)
				if !ok || err != nil {
					// TODO: hmm don't know what the best thing to do logging wise
				}
			}
			copied = append(copied, copyPath)
		}
	}

	return copied, nil
}

// CopyPackratTarball sets up the tarball with the proper naming conventions in a temp directory
// setup tarball takes the path of a tarball and returns a path of the new tarball or an error
// the new tarball will have the proper folder name (the package name)
func CopyPackratTarball(pm rcmd.Package, p string, dest string) (bool, error) {
	dir, err := ioutil.TempDir("", pm.Name)
	defer os.RemoveAll(dir) // clean up
	if err != nil {
		return false, err
	}
	err = archiver.TarGz.Open(p, dir)
	if err != nil {
		return false, err
	}
	tardir, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	newDir := filepath.Join(dir, pm.Name)
	err = os.Rename(filepath.Join(
		dir, tardir[0].Name()),
		newDir)
	if err != nil {
		return false, err
	}
	newTar, err := os.Open(fmt.Sprintf("%s/%s_%s.tar.gz", dest, pm.Name, pm.Version))
	if err != nil {
		return false, err
	}
	err = archiver.TarGz.Write(newTar, []string{newDir})
	if err != nil {
		return false, err
	}

	return true, nil
}
