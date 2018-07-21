package tarutils

import (
	"testing"
)

func TestPackageVersion(t *testing.T) {

	var tartests = []struct {
		in       string
		expected string
	}{
		{
			"testdata/test1zeropointzeropointone.tar.gz",
			"0.0.1",
		},
	}
	for _, tt := range tartests {
		actual := PackageVersion(tt.in)
		if actual != tt.expected {
			t.Errorf("GOT: %s, WANT: %s", actual, tt.expected)
		}
	}
}
