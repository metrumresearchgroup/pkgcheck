package rcmd

// Package stores information about the package
type Package struct {
	Name string `json:"name,omitempty"`
	// this should eventually be some semver type or something hopefully
	Version string `json:"version,omitempty"`
}

// RSettings controls settings related to managing libraries
type RSettings struct {
	LibPaths []string `json:"lib_paths,omitempty"`
	Rpath    string   `json:"rpath,omitempty"`
}

// CheckSettings defines settings related to R CMD CHECK
type CheckSettings struct {
	TarPath   string `json:"tar_path,omitempty"`
	OutputDir string `json:"output_dir,omitempty"`
	Vanilla   bool   `json:"vanilla,omitempty"`
	Cran      bool   `json:"cran,omitempty"`
}
