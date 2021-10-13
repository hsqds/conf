package conf

// LoadResult represents intermediate form for data
// loaded from source. May contain loading error
type LoadResult struct {
	SourceID string
	Config   Config
	Err      error
	Service  string
	Priority int
}

// QueueItem represents
type QueueItem struct {
	value    *LoadResult
	priority int
	index    int
}

// Value
func (qi *QueueItem) Value() *LoadResult {
	return qi.value
}

// NewItem
func NewItem(r *LoadResult) *QueueItem {
	return &QueueItem{
		value:    r,
		priority: r.Priority,
	}
}

type LoadResultsPriorityQueue []*QueueItem

// Len
func (pq LoadResultsPriorityQueue) Len() int {
	return len(pq)
}

// Less
func (pq LoadResultsPriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

// Swap
func (pq LoadResultsPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push
func (pq *LoadResultsPriorityQueue) Push(i interface{}) {
	n := len(*pq)

	lr, _ := i.(*QueueItem)
	lr.index = n
	*pq = append(*pq, lr)
}

// Pop
func (pq *LoadResultsPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]

	return item
}
