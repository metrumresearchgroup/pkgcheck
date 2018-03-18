package rcmd

import "fmt"

// CmdArgs returns a string representation of the command flags associated with
// CheckSettings
func (cs CheckSettings) CmdFlags() []string {
	// by default there is going to be some convention about
	// how much to expose from R CMD check, as this tool has a different
	// scope
	output := []string{
		"--no-manual",
		"--no-build-vignettes",
	}

	if cs.Vanilla {
		output = append(output, "--vanilla")
	}
	if cs.Cran {
		output = append(output, "--as-cran")
	}
	if cs.OutputDir != "" {
		output = append(output, fmt.Sprintf("--output=%s", cs.OutputDir))
	}
	return output
}
