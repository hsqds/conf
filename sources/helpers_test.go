package sources_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf/sources"
)

// TestToCamelCase
func TestToCamelCase(t *testing.T) {
	t.Parallel()

	const (
		delimiter = "_"
	)

	cc := []struct {
		in  string
		out string
	}{
		{"OneWord", "oneword"},
		{"UPPER_CASE", "upperCase"},
		{"lower_case", "lowerCase"},
		{"many_seg_ments_more_and_more", "manySegMentsMoreAndMore"},
	}

	for i := range cc {
		c := cc[i]
		t.Run(c.in, func(t *testing.T) {
			t.Parallel()

			r := sources.ToCamelCase(c.in, delimiter)
			assert.Equal(t, c.out, r)
		})
	}
}

// TestUniqueStrings
func TestUniqueStrings(t *testing.T) {
	t.Parallel()

	input := []string{
		"string1",
		"string2",
		"string1",
		"string2",
		"string3",
		"String1",
		"STRing1",
		"STRING1",
	}

	exp := []string{
		"string1", "string2", "string3",
	}

	ustrs := sources.UniqueStrings(input)

	assert.Equal(t, exp, ustrs)
}
