package tarutils

import (
	"path/filepath"
	"strings"
	"testing"
)

func BenchmarkList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ListTars("/Users/devinp/Downloads/pkglock_snapshot/pkglib/packrat/src")
	}
}

func BenchmarkListAndParseNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tars := ListTars("/Users/devinp/Downloads/pkglock_snapshot/pkglib/packrat/src")
		for _, t := range tars {
			PackageInfo(t)
		}
	}
}
func BenchmarkListAndParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tars := ListTars("/Users/devinp/Downloads/pkglock_snapshot/pkglib/packrat/src")
		for _, t := range tars {
			tarball := filepath.Base(t)
			tarball = strings.TrimSuffix(tarball, ".tar.gz")
			// package tarball stored as <package>_<version>.tar.gz
			packageVersion := strings.SplitN(tarball, "_", 2)
			if len(packageVersion) < 2 {
				PackageInfo(t)
			}
		}
	}
}
