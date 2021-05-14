package sources

import "fmt"

// ServiceConfigError represents.
type ServiceConfigError struct {
	ServiceName string
	SourceID    string
}

// Error.
func (e ServiceConfigError) Error() string {
	return fmt.Sprintf("source %q has no config for service %q", e.SourceID, e.ServiceName)
}
