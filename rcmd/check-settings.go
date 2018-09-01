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

// ShouldCheck returns whether a package should be checked given the filterlist type
// if no filter map is provided, eg. FilterMap{} the default is true
func ShouldCheck(name string, fm FilterMap) bool {
	if name == "" {
		return false
	}
	// if no filtermap set default is to check
	if fm.Type == "" {
		return true
	}
	_, ok := fm.Map[name]
	if fm.Type == "whitelist" {
		return ok
	}
	// if blacklist should not be checked if present
	return !ok
}
