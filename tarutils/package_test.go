package tarutils

import (
	"testing"

	"github.com/r-infra/pkgcheck/rcmd"
)

func TestPackageVersion(t *testing.T) {

	var tartests = []struct {
		in       string
		expected rcmd.Package
	}{
		{
			"testdata/test1zeropointzeropointone.tar.gz",
			rcmd.Package{Name: "test1", Version: "0.0.1"},
		},
		{
			"testdata/pillar/7582a75ff83defed972b348d48b479b8be087f9f.tar.gz",
			rcmd.Package{Name: "pillar", Version: "1.3.0.9000"},
		},
	}
	for _, tt := range tartests {
		actual := PackageInfo(tt.in)
		if actual != tt.expected {
			t.Errorf("GOT: %s, WANT: %s", actual, tt.expected)
		}
	}
}
