package conf

import "fmt"

type ServiceConfigStorageError struct {
	ServiceName string
}

// Error.
func (e ServiceConfigStorageError) Error() string {
	return fmt.Sprintf("could not get %q service config from storage", e.ServiceName)
}

type SourceUniquenessError struct {
	SourceID string
}

// Error.
func (e SourceUniquenessError) Error() string {
	return fmt.Sprintf("source id (%q) is not unique in storage", e.SourceID)
}

//  represents
type SourceStorageError struct {
	SourceID string
}

// Error.
func (e SourceStorageError) Error() string {
	return fmt.Sprintf("source storage has no source %q", e.SourceID)
}

// LoadError represents
type LoadError struct {
	SourceID string
	Service  string
	Err      error
}

// Error
func (e LoadError) Error() string {
	return fmt.Sprintf(
		"could not load service (%q) config from source (%q): %s",
		e.Service, e.SourceID, e.Err,
	)
}

// Unwrap
func (e LoadError) Unwrap() error {
	return e.Err
}
