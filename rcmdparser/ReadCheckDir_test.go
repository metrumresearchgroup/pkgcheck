package rcmdparser

import (
	"reflect"
	"testing"

	"github.com/dpastoor/goutils"

	"github.com/spf13/afero"
)

func TestReadCheckDir(t *testing.T) {
	testFS := afero.NewMemMapFs()
	// tests dir but no testthat
	testFS.MkdirAll("noTestThat/tests", 0755)
	goutils.WriteLinesFS(testFS, []string{"log"}, "noTestThat/00check.log")
	goutils.WriteLinesFS(testFS, []string{"install"}, "noTestThat/00install.out")
	var cdtests = []struct {
		in       string
		expected CheckOutput
	}{
		{
			"noTestThat",
			CheckOutput{
				TestOutput{true, false, nil},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
	}
	for _, tt := range cdtests {
		actual, _ := ReadCheckDir(testFS, tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("GOT: %v, WANT: %v", actual, tt.expected)
		}
	}
}
