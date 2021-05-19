package sources_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/hsqds/conf"
	"github.com/hsqds/conf/sources"
	"github.com/stretchr/testify/assert"
)

// TestJSONSource
func TestJSONSource(t *testing.T) {
	t.Parallel()

	const (
		priority = 50
		svc1     = "service1"
		svc2     = "service2"
	)

	t.Run("ID", func(t *testing.T) {
		t.Parallel()

		s := sources.NewJSONSource(priority, nil)
		assert.Contains(t, s.ID(), "json")
	})

	t.Run("Priority", func(t *testing.T) {
		t.Parallel()

		s := sources.NewJSONSource(priority, nil)
		assert.Equal(t, priority, s.Priority())
	})

	t.Run("Load/ServiceConfig", func(t *testing.T) {
		t.Parallel()

		var (
			svc1ConfData = map[string]string{
				"key1": "value1",
				"key2": "value2",
			}
			svc1Conf = conf.NewMapConfig(svc1ConfData)

			cfg = map[string]map[string]string{
				svc1: svc1ConfData,
			}

			jsonCfg, _ = json.Marshal(cfg)

			newTestJSONCfgSrc = func(b []byte) *sources.JSONSource {
				dataSrc := ioutil.NopCloser(bytes.NewBuffer(b))
				s := sources.NewJSONSource(priority, dataSrc)

				return s
			}

			newPositiveSrcReader = func(data []byte) sources.SrcReader {
				return func(r io.Reader) ([]byte, error) {
					return data, nil
				}
			}

			errSrcReader = func(r io.Reader) ([]byte, error) {
				return nil, errors.New("fail")
			}

			errJSONUnmarshaler = func([]byte, interface{}) error {
				return errors.New("fail")
			}
		)

		t.Run("ServiceConfig should return correct config", func(t *testing.T) {
			t.Parallel()

			s := newTestJSONCfgSrc(jsonCfg)
			s.SetReader(newPositiveSrcReader(jsonCfg))

			err := s.Load(context.Background(), []string{svc1, svc2})
			assert.Nil(t, err)
			r, err := s.ServiceConfig(svc1)

			assert.Nil(t, err)
			assert.Equal(t, svc1Conf, r)
		})

		t.Run("ServiceConfig should return ServiceConfigError when service config is not set", func(t *testing.T) {
			t.Parallel()

			s := newTestJSONCfgSrc(jsonCfg)

			_, err := s.ServiceConfig(svc1)
			assert.IsType(t, sources.ServiceConfigError{}, err)
		})

		t.Run("Load should return error when config reading failed", func(t *testing.T) {
			t.Parallel()

			s := newTestJSONCfgSrc(jsonCfg)
			s.SetReader(errSrcReader)

			err := s.Load(context.Background(), []string{svc1})
			assert.Error(t, err)
		})

		t.Run("Load should return error when unmarshalling failed", func(t *testing.T) {
			t.Parallel()

			s := newTestJSONCfgSrc(jsonCfg)
			s.SetReader(newPositiveSrcReader(jsonCfg))
			s.SetJSONUnmarshaler(errJSONUnmarshaler)

			err := s.Load(context.Background(), []string{svc1})
			assert.Error(t, err)
		})
	})
}
