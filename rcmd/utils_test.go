package rcmd

import (
	"reflect"
	"testing"
)

func makeInMap(wl []string, bl []string) map[string][]string {
	output := make(map[string][]string)
	output["whitelist"] = wl
	output["blacklist"] = bl
	return output
}
func makeOutputMap(pkgs []string, t string) FilterMap {
	output := make(map[string]bool)
	for _, pkg := range pkgs {
		output[pkg] = true
	}
	return FilterMap{
		Type: t,
		Map:  output,
	}
}
func TestFilterMap(t *testing.T) {
	var fmTests = []struct {
		in       map[string][]string
		expected FilterMap
	}{
		{
			in:       makeInMap([]string{"dplyr", "ggplot2"}, []string{}),
			expected: makeOutputMap([]string{"dplyr", "ggplot2"}, "whitelist"),
		},
		{
			in:       makeInMap([]string{"dplyr", "ggplot2"}, []string{"tidyverse"}),
			expected: makeOutputMap([]string{"dplyr", "ggplot2"}, "whitelist"),
		},
		{
			in:       makeInMap([]string{}, []string{"tidyverse"}),
			expected: makeOutputMap([]string{"tidyverse"}, "blacklist"),
		},
	}
	for _, tt := range fmTests {
		fm := CreateFilterMap(tt.in["whitelist"], tt.in["blacklist"])
		eq := reflect.DeepEqual(fm, tt.expected)
		if !eq {
			t.Errorf("filter Maps not equal, got: %v, expected %v", fm, tt.expected)
		}
	}
}
