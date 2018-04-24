package rcmdparser

import (
	"github.com/spf13/afero"
)

// NewCheck creates a new CheckOutput Object
// fs and check directory
func NewCheck(fs afero.Fs, cd string) (LogResults, error) {
	cr, err := ReadCheckDir(fs, cd)
	if err != nil {
		return LogResults{}, err
	}
	return cr.Parse(), nil
}

// Parse output to LogResults
func (c CheckOutput) Parse() LogResults {
	lr := LogResults{
		Checks: ParseCheckLog(c.Check),
	}
	if c.Test.Testthat {
		lr.Tests = ParseTestLog(c.Test.Results)
	}
	return lr
}
