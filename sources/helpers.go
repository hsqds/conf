package sources

import (
	"strings"
)

// toCamelCase.
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

// uniqueStrings returns slice of strings containing only unique strings
// !! Case-insensitive
func uniqueStrings(strs []string) []string {
	m := make(map[string]struct{}, len(strs))
	ustrs := make([]string, 0, len(strs))

	var s string
	for i := range strs {
		s = strings.ToUpper(strs[i])
		if _, ok := m[s]; ok {
			continue
		}

		m[s] = struct{}{}

		ustrs = append(ustrs, strs[i])
	}

	return ustrs
}
