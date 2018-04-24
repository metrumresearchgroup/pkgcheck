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

// LogEntries are the parsed results from the check log
type LogEntries struct {
	Errors   []string
	Warnings []string
	Notes    []string
}

// LogResults is a struct of the R CMD check results
type LogResults struct {
	Checks LogEntries
	Tests  TestResults
}
