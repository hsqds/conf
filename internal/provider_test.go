package internal_test

import "testing"

// TestProvider
func TestProvider(t *testing.T) {
	t.Run("NewProvider should create new provider instance", func(t *testing.T) {
		NewProvider()
	})
}
