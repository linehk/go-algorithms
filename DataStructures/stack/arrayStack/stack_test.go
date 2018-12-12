package arrayStack

import (
	"testing"

	"GoAlgorithms/utils"
)

func TestStack(t *testing.T) {
	tests := []struct {
		seq  []interface{}
		want []interface{}
	}{
		{[]interface{}{0, 1, 2, 3}, []interface{}{3, 2, 1, 0}},
	}
	for i, tt := range tests {
		s := New(len(tt.seq))
		for _, v := range tt.seq {
			if err := s.Push(v); err != nil {
				t.Error(err)
			}
		}

		got := make([]interface{}, 0)
		for range tt.seq {
			v, err := s.Pop()
			if err != nil {
				t.Error(err)
			}
			got = append(got, v)
		}

		if !utils.IsSameSlice(got, tt.want) {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
