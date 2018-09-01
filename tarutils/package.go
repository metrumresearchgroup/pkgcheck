package tarutils

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/metrumresearchgroup/pkgcheck/rcmd"
)

// PackageVersion returns the package version from the Version field of the DESCRIPTION file
// This can be useful if a package does not follow the package_version.tar.gz convention
func PackageInfo(tarpath string) rcmd.Package {
	output := rcmd.Package{}
	file, err := os.Open(tarpath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	archive, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	tr := tar.NewReader(archive)

FILELOOP:
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if hdr.FileInfo().Name() == "DESCRIPTION" {
			// create a Scanner for reading line by line
			s := bufio.NewScanner(tr)

			// line reading loop
			for s.Scan() {

				// read the current last read line of text
				l := s.Text()
				if strings.HasPrefix(l, "Version:") {
					output.Version = strings.TrimPrefix(l, "Version: ")
				}
				if strings.HasPrefix(l, "Package:") {
					output.Name = strings.TrimPrefix(l, "Package: ")
				}
			}
			break FILELOOP
		}
	}
	return output
}
