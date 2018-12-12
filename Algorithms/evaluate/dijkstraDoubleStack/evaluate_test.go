package dijkstraDoubleStack

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"(1+1)", 2},
	}
	for i, tt := range tests {
		if got := Eval(tt.s); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
