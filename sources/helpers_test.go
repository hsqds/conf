package sources_test

import (
	"strings"
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

	t.Run("should convert snake_case to camelCase", func(t *testing.T) {
		t.Parallel()

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
	})

	t.Run("should also works with different delimiters", func(t *testing.T) {
		t.Parallel()

		const (
			tplDelim = "_"
			tpl      = "Here_comes_different_separators"
			exp      = "hereComesDifferentSeparators"
		)

		delimiters := []string{"*", "#", "##", "%%%", "&", "(", ")", "\"\""}

		for _, d := range delimiters {
			t.Run(d, func(t *testing.T) {
				t.Parallel()

				r := sources.ToCamelCase(strings.ReplaceAll(tpl, tplDelim, d), d)
				assert.Equal(t, exp, r)
			})
		}
	})
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

// TestSieveServiceConfig
func TestSieveServiceConfig(t *testing.T) {
	t.Parallel()

	const (
		svc        = "serviceName"
		prefix     = `\+\+`
		delimiter  = `\*`
		assignment = `:=`
	)

	input := []string{
		"++serviceName*key1:=value1",
		"--serviceName*key2:=value2",
		"++servicename*key3 := value3",
		"++servicename*key*name4:=value4",
	}

	exp := map[string]string{
		"key1":     "value1",
		"keyName4": "value4",
	}

	r := sources.SieveServiceConfig(svc, prefix, delimiter, assignment, input)

	assert.Equal(t, exp, r)
}
