package viralloadhistory

import "sync"

type SnapshotPolicy struct {
	Mode       string
	CacheReads bool
}

type History struct {
	mu     sync.RWMutex
	values []int
	policy SnapshotPolicy
}

func NewHistory(values []int, policy SnapshotPolicy) *History {
	return &History{values: append([]int(nil), values...), policy: policy}
}

func (h *History) Append(value int) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.values = append(h.values, value)
}

func (h *History) Values() []int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.policy.Mode == "fast" && h.policy.CacheReads {
		return h.values
	}
	return append([]int(nil), h.values...)
}
