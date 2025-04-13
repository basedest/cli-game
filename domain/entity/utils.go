package entity

import "strings"

// joinStrings joins string slices with the provided separator
func joinStrings(items []string, separator string) string {
	return strings.Join(items, separator)
} 