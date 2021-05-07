package stubs

import (
	"fmt"

	"github.com/google/uuid"
)

// TestConfig represents `provider.Config` interface
// implementation for testing
type TestConfig struct {
	value string
}

// NewTestConfig
func NewTestConfig() TestConfig {
	return TestConfig{
		value: fmt.Sprintf("test config id: %s", uuid.NewString()),
	}
}

// Get
func (c *TestConfig) Get(key, defaultValue string) string {
	return c.value
}

// Fmt
func (c *TestConfig) Fmt(pattern string) (string, error) {
	return "", nil
}
