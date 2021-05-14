package sources_test

import (
	"strings"
	"testing"

	"github.com/hsqds/conf/sources"
)

// TestServiceConfigError
func TestServiceConfigError(t *testing.T) {
	t.Run("error message should contain source ID and service name", func(t *testing.T) {
		const (
			svcName = "testService1231"
			srcID   = "lskdf9832u982"
		)

		e := sources.ServiceConfigError{ServiceName: svcName, SourceID: srcID}
		msg := e.Error()

		if !strings.Contains(msg, svcName) {
			t.Errorf("message should contain service name %q", svcName)
		}

		if !strings.Contains(msg, srcID) {
			t.Errorf("message should contain source id %q", srcID)
		}
	})
}
