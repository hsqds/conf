package sources

import "context"

// MapSource represents
type MapSource map[string]string

// Get provide value by key
func (s MapSource) Get(ctx context.Context, key string) (string, error) {
	value, ok := s[key]
	if !ok {
		return "", ErrNoSuchKey
	}

	return value, nil
}

// GetAll
func (s MapSource) GetAll(ctx context.Context) (map[string]string, error) {
	return s, nil
}
