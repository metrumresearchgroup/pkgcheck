package rcmd

// CreateFilterMap creates a filter map based on a whitelist and blacklist
// defaults to prioritizing a whitelist, such that if any present
// then the blacklist is ignored
func CreateFilterMap(wl []string, bl []string) FilterMap {
	filterMap := make(map[string]bool)
	filterType := "blacklist"
	if len(wl) > 0 {
		filterType = "whitelist"
		for _, pkg := range wl {
			filterMap[pkg] = true
		}
	} else {
		for _, pkg := range bl {
			filterMap[pkg] = true
		}
	}
	return FilterMap{
		Type: filterType,
		Map:  filterMap,
	}
}
