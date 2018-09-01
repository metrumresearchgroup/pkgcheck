package main

import (
	"github.com/metrumresearchgroup/pkgcheck/rcmd"
	"github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

func main() {
	appFS := afero.NewOsFs()
	lg := logrus.New()
	lg.SetLevel(logrus.DebugLevel)
	rcmd.RunR(appFS, rcmd.RSettings{
		LibPaths: []string{},
		Rpath:    "R",
	},
		"/Users/devin/clients/amgen/pdms",
		lg)

}
