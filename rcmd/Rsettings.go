package rcmd

import (
	"strings"
)

// R provides a cleaned path to the R executable
func (rs RSettings) R() string {
	// TODO: check if this could have problems with trailing slash on windows
	// TODO: better to use something like filepath.clean? would that sanitize better?

	// Need to trim trailing slash as will form the R CMD syntax
	// eg /path/to/R CMD, so can't have /path/to/R/ CMD
	return strings.TrimSuffix(rs.Rpath, "/")
}
