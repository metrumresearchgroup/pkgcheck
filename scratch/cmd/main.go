package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/afero"
)

func main() {
	appFS := afero.NewOsFs()

	file, _ := afero.ReadFile(appFS, "../../parser/testdata/testwarning.Rcheck/00check.log")
	splitFile := bytes.Split(file, []byte("* "))

	for _, ent := range splitFile {
		fmt.Println(string(ent))
	}
}
