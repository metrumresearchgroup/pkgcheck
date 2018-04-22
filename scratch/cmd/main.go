package main

import (
	"fmt"

	"github.com/dpastoor/pkgcheck/rcmdparser"
	"github.com/spf13/afero"
)

func main() {
	appFS := afero.NewOsFs()

	checkDir := "../../rcmdparser/testdata/testerror.Rcheck"
	fmt.Println(rcmdparser.ReadCheckDir(appFS, checkDir))
}
