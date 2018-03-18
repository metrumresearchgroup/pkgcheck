package rcmd

import "testing"

var cstests = []struct {
	in       CheckSettings
	expected []string
}{
	{CheckSettings{},
		[]string{
			"--no-manual",
			"--no-build-vignettes",
		},
	},
}

func TestCheckSettings(t *testing.T) {
	for _, tt := range cstests {
		for i, actual := range tt.in.CmdFlags() {

			if actual != tt.expected[i] {
				t.Errorf("GOT: %s, WANT: %s", actual, tt.expected[i])
			}
		}
	}
}
