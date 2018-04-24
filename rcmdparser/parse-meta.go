package rcmdparser

import (
	"bytes"
)

// Parse consumes a log entry and updates the check metadata if relevant
func (cm *CheckMeta) Parse(ent []byte) {
	switch {
	case bytes.Contains(ent, []byte("log directory")):
		cm.LogDir = extractBetweenQuotes(ent)
	case bytes.Contains(ent, []byte("R version")):
		cm.Rversion = parseRVersion(ent)
	case bytes.Contains(ent, []byte("platform:")):
		cm.Platform = string(
			bytes.Replace(ent,
				[]byte("using platform: "),
				[]byte(""), 1),
		)
	case bytes.Contains(ent, []byte("options")):
		cm.Options = extractBetweenQuotes(ent)
	case bytes.Contains(ent, []byte("this is package")):
		cm.Package = ""
		cm.PackageVersion = ""
	default:
	}
}

func extractBetweenQuotes(ent []byte) string {
	sb := bytes.Index(ent, []byte("‘"))
	eb := bytes.Index(ent, []byte("’"))
	if sb == -1 || eb == -1 {
		// didn't parse correctly, return whole entry
		return string(ent)
	}
	// when trying to just clip bytes eg sb+1:eb
	// was getting weird printing artifact, so the index
	// trimming to remote
	return string(bytes.TrimPrefix(ent[sb:eb], []byte("‘")))
}

func parseRVersion(ent []byte) string {
	return ""
}
