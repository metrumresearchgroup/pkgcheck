package tarutils

import "github.com/metrumresearchgroup/pkgcheck/rcmd"

// PackageTarball gives the package details, and whether the PackageTarball
// is stored as a hashed package. In the case of github packages, packrat
// stores the package via the format <org>-<pkgname>-<hash>. This is incompatible
// with RCMD CHECK, which requires it to just be the package name
type PackageTarball struct {
	Details       rcmd.Package
	PackratHashed bool
}
