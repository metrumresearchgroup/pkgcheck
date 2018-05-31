package rcmdparser

import (
	"github.com/spf13/afero"
)

// NewCheck creates a new CheckOutput Object
// fs and check directory
func NewCheck(fs afero.Fs, cd string) (CheckResults, error) {
	cr, err := ReadCheckDir(fs, cd)
	if err != nil {
		return CheckResults{}, err
	}
	return cr.Parse(), nil
}

// Parse output to LogResults
func (c CheckData) Parse() CheckResults {
	lr := CheckResults{
		Checks: ParseCheckLog(c.Check),
	}
	if c.Test.Testthat {
		lr.Tests = ParseTestLog(c.Test.Results)
	}
	return lr
}
