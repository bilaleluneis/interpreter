package parser

import "strings"

// normalize removes all tab characters from a string.
func normalize(s string) string {
	return strings.ReplaceAll(s, "\t", "")
}
