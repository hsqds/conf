package sources

import (
	"strings"
)

// toCamelCase
func toCamelCase(key, delimiter string) string {
	const (
		minDelimCount = 1
	)

	if strings.Count(key, delimiter) < minDelimCount {
		return strings.ToLower(key)
	}

	segments := strings.Split(key, delimiter)
	segments[0] = strings.ToLower(segments[0])

	for i := 1; i < len(segments); i++ {
		segments[i] = strings.Title(strings.ToLower(segments[i]))
	}

	return strings.Join(segments, "")
}
