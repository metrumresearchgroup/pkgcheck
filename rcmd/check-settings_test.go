package rcmd

import (
	"testing"
)

func TestCmdFlags(t *testing.T) {

	var cstests = []struct {
		in       CheckSettings
		expected []string
	}{
		{
			CheckSettings{
				TarPath: "path/to/dplyr_0.7.4.tar.gz",
			},
			[]string{
				"--no-manual",
				"--no-build-vignettes",
			},
		},
	}
	for _, tt := range cstests {
		for i, actual := range tt.in.CmdFlags() {

			if actual != tt.expected[i] {
				t.Errorf("GOT: %s, WANT: %s", actual, tt.expected[i])
			}
		}
	}
}

func TestPackage(t *testing.T) {

	var packages = []struct {
		in       CheckSettings
		expected Package
	}{
		{
			CheckSettings{
				TarPath: "dplyr_0.7.4.tar.gz",
			},
			Package{
				Name:    "dplyr",
				Version: "0.7.4",
			},
		},
		{
			CheckSettings{
				TarPath: "path/to/dplyr_0.7.4.tar.gz",
			},
			Package{
				Name:    "dplyr",
				Version: "0.7.4",
			},
		},
	}
	for _, tt := range packages {
		actual := tt.in.Package()
		if actual != tt.expected {
			t.Errorf("GOT: %s, WANT: %s", actual, tt.expected)
		}
	}
}
