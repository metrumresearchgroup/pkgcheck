package rcmdparser

// TestData represents output of tests and whether it uses testthat
type TestData struct {
	HasTests bool
	Testthat bool
	Results  []byte
}

// CheckData represents key elements of a R CMD check output directory
type CheckData struct {
	Test    TestData
	Check   []byte
	Install []byte
}

// TestResults is the results from testthat
type TestResults struct {
	Ok        int
	Skipped   int
	Failed    int
	Output    string
	Available bool
}

// CheckMeta stores metadata about the RCMDCHECK
type CheckMeta struct {
	LogDir         string
	Rversion       string
	Platform       string
	Options        string
	Package        string
	PackageVersion string
}

// LogEntries are the parsed results from the check log
type LogEntries struct {
	Meta     CheckMeta
	Errors   []string
	Warnings []string
	Notes    []string
}

// CheckResults is a struct of the R CMD check results
type CheckResults struct {
	Checks LogEntries
	Tests  TestResults
}
