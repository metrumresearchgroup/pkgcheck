package main

import (
	"github.com/sirupsen/logrus"

	"github.com/dpastoor/pkgcheck/rcmdparser"
	"github.com/spf13/afero"
)

func main() {
	appFS := afero.NewOsFs()
	lg := logrus.New()
	checkDir := "../../rcmdparser/testdata/testerror.Rcheck"
	results, err := rcmdparser.NewCheck(appFS, checkDir)
	if err != nil {
		panic(err)
	}
	results.Log(lg)

}
