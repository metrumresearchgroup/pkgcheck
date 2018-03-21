package rcmd

// Package stores information about the package
type Package struct {
	Name string
	// this should eventually be some semver type or something hopefully
	Version string
}

// RSettings controls settings related to managing libraries
type RSettings struct {
	LibPaths []string
	Rpath    string
	LibDirs  []string
}

// CheckSettings defines settings related to R CMD CHECK
type CheckSettings struct {
	TarPath   string
	OutputDir string
	Vanilla   bool
	Cran      bool
}
