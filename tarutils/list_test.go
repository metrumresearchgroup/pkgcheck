package tarutils

import (
	"testing"

	"github.com/r-infra/pkgcheck/rcmd"

	"github.com/spf13/afero"
)

func TestListingTars(t *testing.T) {

	// this is still a bit mysterious, the walk function used in ListTars should be returning
	// the absolute paths, per empirical testing, however within the testdata dir, it instead
	// returns a relative path to testdata. This is a pleasant side effect, given it would be
	// annoying to normalize all paths for different systems when testing since the
	// resulting paths are coded into the expected results.
	// However, this is somewhat concerning as I am unsure if there are other instances in which
	// a non-absolute path is returned.
	var tartests = []struct {
		in       string
		expected []string
	}{
		{
			"testdata",
			[]string{"testdata/pillar/7582a75ff83defed972b348d48b479b8be087f9f.tar.gz", "testdata/test1zeropointzeropointone.tar.gz"},
		},
	}
	for _, tt := range tartests {
		actual := ListTars(tt.in)
		if len(actual) != len(tt.expected) {
			t.Fatalf("Mismatch in length, GOT: %s, WANT: %s", actual, tt.expected)
		}
		for i, p := range actual {
			if p != tt.expected[i] {
				t.Errorf("GOT: %s, WANT: %s", actual, tt.expected)
			}
		}
	}
}

func TestCopyTars(t *testing.T) {

	var tartests = []struct {
		in       string
		expected []string
	}{
		{
			"testdata",
			[]string{"tarballs/pillar_1.3.0.9000.tar.gz", "tarballs/test1_0.0.1.tar.gz"},
		},
	}
	fs := afero.NewMemMapFs()
	for _, tt := range tartests {
		actual, err := CopyPkgTars(fs, tt.in, "tarballs", rcmd.FilterMap{})
		if err != nil {
			t.Fatal(err)
		}
		for i, p := range actual {
			if p != tt.expected[i] {
				t.Errorf("GOT: %s, WANT: %s", actual, tt.expected)
			}
		}
	}
}
