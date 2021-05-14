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
