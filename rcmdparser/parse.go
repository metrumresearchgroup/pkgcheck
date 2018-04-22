package rcmdparser

import (
	"bytes"
)

// ParseCheckLog parses the check log
func ParseCheckLog(e []byte) CheckResults {
	splitOutput := bytes.Split(e, []byte("* "))
	var errors []string
	var notes []string
	var warnings []string

	for _, ent := range splitOutput {
		switch {
		case bytes.Contains(ent, []byte("... NOTE")):
			notes = append(notes, string(ent))
		case bytes.Contains(ent, []byte("... ERROR")):
			errors = append(errors, string(ent))
		case bytes.Contains(ent, []byte("... WARNING")):
			warnings = append(warnings, string(ent))
		default:
			continue
		}
	}
	return CheckResults{
		Errors:   errors,
		Notes:    notes,
		Warnings: warnings,
	}
}
