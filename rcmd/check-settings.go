package rcmd

import (
	"fmt"
	"path/filepath"
	"strings"
)

// CmdFlags returns a string representation of the command flags associated with
// CheckSettings
func (cs CheckSettings) CmdFlags() []string {
	// by default there is going to be some convention about
	// how much to expose from R CMD check, as this tool has a different
	// scope
	output := []string{
		"--no-manual",
		"--no-build-vignettes",
	}

	if cs.Cran {
		output = append(output, "--as-cran")
	}
	if cs.OutputDir != "" {
		output = append(output, fmt.Sprintf("--output=%s", cs.OutputDir))
	}
	return output
}

// Package returns information about the package information
func (cs CheckSettings) Package() Package {
	if cs.TarPath == "" {
		return Package{}
	}
	tarball := filepath.Base(cs.TarPath)
	tarball = strings.TrimSuffix(tarball, ".tar.gz")
	// package tarball stored as <package>_<version>.tar.gz
	packageVersion := strings.SplitN(tarball, "_", 2)
	return Package{
		Name:    packageVersion[0],
		Version: packageVersion[1],
	}
}

// InFilterList checks if a package is in a given whitelist or blacklist
func InFilterList(cs CheckSettings, fl map[string]bool) bool {
	pkgName := cs.Package().Name
	_, ok := fl[pkgName]
	return ok
}
