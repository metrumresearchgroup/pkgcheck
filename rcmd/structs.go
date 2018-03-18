package rcmd

// Rsettings controls settings related to managing libraries
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
