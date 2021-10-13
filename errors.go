package conf

import "fmt"

type ErrServiceConfigNotFound struct {
	ServiceName string
}

// Error.
func (e ErrServiceConfigNotFound) Error() string {
	return fmt.Sprintf("could not get %q service config from storage", e.ServiceName)
}

type ErrSourceIsNotUnique struct {
	SourceID string
}

// Error.
func (e ErrSourceIsNotUnique) Error() string {
	return fmt.Sprintf("source id (%q) is not unique in storage", e.SourceID)
}

//  represents
type ErrSourceNotFound struct {
	SourceID string
}

// Error.
func (e ErrSourceNotFound) Error() string {
	return fmt.Sprintf("source storage has no source %q", e.SourceID)
}
