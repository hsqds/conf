package sources_test

import (
	"context"
	"testing"

	"github.com/go-test/deep"

	"github.com/hate-squids/config-provider/internal/sources"
)

// TestMapSource
func TestMapSource(t *testing.T) {
	t.Run("should return correct value by key", func(t *testing.T) {
		key := "key"
		value := "some text"
		s := sources.MapSource(map[string]string{
			key:            value,
			"one more key": "birds",
			"another key":  "cats",
		})

		v, err := s.Get(context.Background(), key)
		if err != nil {
			t.Fail()
		}

		if v != value {
			t.Error("source could return wrong value")
		}
	})

	t.Run("should return all values", func(t *testing.T) {
		d := map[string]string{
			"key1": "value",
			"key2": "value",
		}
		s := sources.MapSource(d)

		vs, err := s.GetAll(context.Background())
		if err != nil {
			t.Fail()
		}

		if diff := deep.Equal(d, vs); diff != nil {
			t.Error(diff)
		}
	})

	t.Run("should return ErrNoSuchKey if key is incorrect", func(t *testing.T) {
		s := sources.MapSource(map[string]string{
			"validKey": "value",
		})

		_, err := s.Get(context.Background(), "invalidKey")
		if err != sources.ErrNoSuchKey {
			t.Fail()
		}
	})
}
