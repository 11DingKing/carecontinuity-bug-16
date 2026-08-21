package viralloadhistory_test

import (
	"reflect"
	"testing"

	"carecontinuity/internal/continuity/viralloadhistory"
)

func TestViralLoadHistoryPublicBehavior(t *testing.T) {
	coordinator := viralloadhistory.NewCoordinator([]int{17, 29, 43})
	first := coordinator.Snapshot()
	first[0], first[2] = 99, 1
	if got, want := coordinator.Snapshot(), []int{17, 29, 43}; !reflect.DeepEqual(got, want) {
		t.Fatalf("stored history was changed through returned snapshot: got %v want %v", got, want)
	}
	coordinator.Record(61)
	if got := coordinator.Snapshot(); !reflect.DeepEqual(got, []int{17, 29, 43, 61}) {
		t.Fatalf("append path returned %v", got)
	}
}
