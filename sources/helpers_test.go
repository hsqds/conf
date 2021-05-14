package sources_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hsqds/conf/sources"
)

// TestToCamelCase
func TestToCamelCase(t *testing.T) {
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

	for _, c := range cc {
		t.Run("format to camelCase", func(t *testing.T) {
			r := sources.ToCamelCase(c.in, delimiter)
			if r != c.out {
				t.Errorf("expect %q got %q", c.out, r)
			}
		})
	}
}

// TestUniqueStrings
func TestUniqueStrings(t *testing.T) {
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
