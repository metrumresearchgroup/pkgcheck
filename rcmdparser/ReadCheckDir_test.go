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

	// with testthat
	testFS.MkdirAll("WithTestThat/tests", 0755)
	goutils.WriteLinesFS(testFS, []string{"log"}, "WithTestThat/00check.log")
	goutils.WriteLinesFS(testFS, []string{"install"}, "WithTestThat/00install.out")
	goutils.WriteLinesFS(testFS, []string{"tests"}, "WithTestThat/tests/testthat.Rout")

	// Failed Testthat
	testFS.MkdirAll("FailedTest/tests", 0755)
	goutils.WriteLinesFS(testFS, []string{"log"}, "FailedTest/00check.log")
	goutils.WriteLinesFS(testFS, []string{"install"}, "FailedTest/00install.out")
	goutils.WriteLinesFS(testFS, []string{"failed-tests"}, "FailedTest/tests/testthat.Rout.fail")
	var cdtests = []struct {
		in       string
		expected CheckData
	}{
		{
			"noTestThat",
			CheckData{
				TestData{true, false, nil},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			"WithTestThat",
			CheckData{
				TestData{true, true, []byte("tests\n")},
				[]byte("log\n"),
				[]byte("install\n"),
			},
		},
		{
			"FailedTest",
			CheckData{
				TestData{true, true, []byte("failed-tests\n")},
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
