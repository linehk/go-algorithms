package linkedListQueue

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
		q := New()

		for _, v := range tt.seq {
			q.Enqueue(v)
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
