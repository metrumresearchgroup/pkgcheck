package rcmdparser

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Log the results to the logger
// Prints at InfoLevel
func (lr CheckResults) Log(lg *logrus.Logger) {
	cr := lr.Checks
	tr := lr.Tests
	lg.Infof("%s v%s CHECK RESULTS: %v ERRORS, %v WARNINGS, %v NOTES",
		lr.Checks.Meta.Package,
		lr.Checks.Meta.PackageVersion,
		len(cr.Errors), len(cr.Warnings), len(cr.Notes),
	)
	if tr.Available {
		lg.Infof("%s v%s TEST RESULTS: %v OK, %v Skipped, %v Failed",
			lr.Checks.Meta.Package,
			lr.Checks.Meta.PackageVersion,
			tr.Ok, tr.Skipped, tr.Failed)
	} else {
		lg.Infof("%s v%s TEST RESULTS: NO TESTS PRESENT",
			lr.Checks.Meta.Package,
			lr.Checks.Meta.PackageVersion)
	}
}

// Print the results to stdout
func (lr CheckResults) Print() {
	cr := lr.Checks
	tr := lr.Tests
	fmt.Println("RCMD CHECK RESULTS: ")
	fmt.Println(fmt.Sprintf("%v ERRORS, %v WARNINGS, %v NOTES",
		len(cr.Errors), len(cr.Warnings), len(cr.Notes)))
	if tr.Available {
		fmt.Println("TEST RESULTS:")
		fmt.Println(fmt.Sprintf("%v OK, %v Skipped, %v Failed",
			tr.Ok, tr.Skipped, tr.Failed))
	} else {
		fmt.Println("No Tests Present")
	}
}
