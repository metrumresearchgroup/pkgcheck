package main

import (
	"fmt"

	"github.com/metrumresearchgroup/pkgcheck/tarutils"
)

func main() {
	// appFS := afero.NewOsFs()
	// lg := logrus.New()
	tars := tarutils.ListTars("/Users/devinp/Downloads/pkglock_snapshot/pkglib/packrat/src")
	for _, t := range tars {
		ti := tarutils.PackageInfo(t)
		fmt.Println(fmt.Sprintf("package: %s, version: %s", ti.Name, ti.Version))
	}
}
