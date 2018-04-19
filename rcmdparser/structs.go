package rcmdparser

// TestOutput represents output of tests and whether it uses testthat
type TestOutput struct {
	HasTests bool
	Testthat bool
	Results  []byte
}

// CheckOutput represents key elements of a R CMD check output directory
type CheckOutput struct {
	Test    TestOutput
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

// LogResults is a struct of the R CMD check results
type LogResults struct {
	Errors   []string
	Warnings []string
	Notes    []string
	Tests    TestResults
}
