package tar

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// PackageVersion returns the package version from the Version field of the DESCRIPTION file
// This can be useful if a package does not follow the package_version.tar.gz convention
func PackageVersion(tarpath string) string {
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
				if strings.HasPrefix(l, "Version") {
					return strings.TrimPrefix(l, "Version: ")
				}
			}
		}
	}
	return ""
}
