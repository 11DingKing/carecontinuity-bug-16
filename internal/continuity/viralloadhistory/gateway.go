package viralloadhistory

type Coordinator struct{ history *History }

func NewCoordinator(values []int) *Coordinator {
	policy := SnapshotPolicy{Mode: "fast", CacheReads: true}
	return &Coordinator{history: NewHistory(values, policy)}
}

func (c *Coordinator) Snapshot() []int  { return c.history.Values() }
func (c *Coordinator) Record(value int) { c.history.Append(value) }
