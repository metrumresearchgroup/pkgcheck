package rcmdparser

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Log the results to the logger
// Prints at InfoLevel
func (lr LogResults) Log(lg *logrus.Logger) {
	cr := lr.Checks
	tr := lr.Tests
	lg.Infoln("RCMD CHECK RESULTS: ")
	lg.Infoln(fmt.Sprintf("%v ERRORS, %v WARNINGS, %v NOTES",
		len(cr.Errors), len(cr.Warnings), len(cr.Notes)))
	if tr.Available {
		lg.Infoln("TEST RESULTS:")
		lg.Infoln(fmt.Sprintf("%v OK, %v Skipped, %v Failed",
			tr.Ok, tr.Skipped, tr.Failed))
	} else {
		lg.Infoln("No Tests Present")
	}
}

// Print the results to stdout
func (lr LogResults) Print() {
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
