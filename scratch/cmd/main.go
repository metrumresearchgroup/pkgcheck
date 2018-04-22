package main

import (
	"fmt"

	"github.com/dpastoor/pkgcheck/rcmdparser"
	"github.com/spf13/afero"
)

func main() {
	appFS := afero.NewOsFs()

	checkDir := "../../rcmdparser/testdata/testwarningerror.Rcheck"
	output, err := rcmdparser.ReadCheckDir(appFS, checkDir)
	if err != nil {
		panic(err)
	}
	checkResults := rcmdparser.ParseCheckLog(output.Check)
	fmt.Println(fmt.Sprintf("%v ERRORS, %v WARNINGS, %v NOTES",
		len(checkResults.Errors), len(checkResults.Warnings), len(checkResults.Notes)))
}
