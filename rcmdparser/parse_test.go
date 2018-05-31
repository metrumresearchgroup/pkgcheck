package rcmdparser

import (
	"reflect"
	"testing"
)

func TestParseTestLog(t *testing.T) {

	var pttests = []struct {
		in       []byte
		expected TestResults
	}{
		{
			[]byte(`
				Stuff above

				> library(testthat)
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 1 SKIPPED: 0 FAILED: 0
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`),
			TestResults{
				Ok:      1,
				Skipped: 0,
				Failed:  0,
				Output: `
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 1 SKIPPED: 0 FAILED: 0
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`,
				Available: true,
			},
		},
		// test double digits
		{
			[]byte(`
				Stuff above

				> library(testthat)
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 12 SKIPPED: 10 FAILED: 11 
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`),
			TestResults{
				Ok:      12,
				Skipped: 10,
				Failed:  11,
				Output: `
				> library(test1)
				> 
				> test_check("test1")
				══ testthat results  ═══════════════════════════════════════════════════════════
				OK: 12 SKIPPED: 10 FAILED: 11 
				> 
				> proc.time()
				   user  system elapsed 
				  0.866   0.075   0.874 
				`,
				Available: true,
			},
		},
		{
			nil,
			TestResults{
				Ok:        0,
				Skipped:   0,
				Failed:    0,
				Available: false,
			},
		},
	}
	for _, tt := range pttests {
		actual := ParseTestLog(tt.in)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("GOT: %v, WANT: %v", actual, tt.expected)
		}
	}
}
