package sources

import (
	"fmt"
	"regexp"
	"strings"
)

// toCamelCase.
func toCamelCase(key, delimiter string) string {
	const (
		minDelimCount = 1
	)

	// delimiter may contain escape-symbols
	delimiter = strings.ReplaceAll(delimiter, "\\", "")

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

// sieveServiceConfig filters `bucket` and returns serviceConfig
// service, prefix, delimiter, assignment - will be embedded in to regexp
// pattern - all special characters should be escaped
func sieveServiceConfig(service, prefix, delimiter, assignment string, bucket []string) map[string]string {
	const (
		reKeyIndex = 2
		reValIndex = 3
		matchesNum = 4
	)

	result := make(map[string]string)
	pattern := fmt.Sprintf(
		`((?i)%s%s%s)([\w%s]+)%s([\w]+)`,
		prefix, service, delimiter, delimiter, assignment,
	)
	svcRe := regexp.MustCompile(pattern)

	for i := range bucket {
		matches := svcRe.FindStringSubmatch(bucket[i])

		if len(matches) < matchesNum {
			continue
		}

		key := toCamelCase(matches[reKeyIndex], delimiter)
		result[key] = matches[reValIndex]
	}

	return result
}
