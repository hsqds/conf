package conf_test

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/hsqds/conf"

	"github.com/stretchr/testify/assert"
)

// TestLoadResultPriorityQueue
func TestLoadResultPriorityQueue(t *testing.T) {
	t.Parallel()

	t.Run("Push/Len", func(t *testing.T) {
		t.Parallel()

		const (
			count = 10
		)

		pq := make(conf.LoadResultsPriorityQueue, 0, count)
		heap.Init(&pq)

		for i := 0; i < count; i++ {
			heap.Push(&pq, conf.NewItem(&conf.LoadResult{
				SourceID: "test-source",
				Priority: i,
				Service:  "test service",
				Config:   conf.NewMapConfig(map[string]string{"val": fmt.Sprintf("%d", i)}),
			}))
		}

		m, _ := heap.Pop(&pq).(*conf.QueueItem)
		p := m.Value().Config.Get("val", "0")

		assert.Equal(t, fmt.Sprintf("%d", count-1), p)
	})
}
