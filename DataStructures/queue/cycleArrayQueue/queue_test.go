package cycleArrayQueue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	tests := []struct {
		seq []interface{}
	}{
		{[]interface{}{0, 1, 2, 3}},
	}
	for i, tt := range tests {
		q := New(len(tt.seq))

		for _, v := range tt.seq {
			if err := q.Enqueue(v); err != nil {
				t.Error(err)
			}
		}

		for _, v := range tt.seq {
			got, err := q.Dequeue()
			if err != nil {
				t.Error(err)
			}

			if got != v {
				t.Errorf("%v. got %v, want %v",
					i, got, v)
			}
		}
	}
}
