package priorityQueue

import (
	"testing"
)

func TestNewByArray(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{0}},
		{[]int{0, 1}},
		{[]int{0, 0}},
		{[]int{0, 1, 2, 3, 4, 5, 6}},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		if got := p.isMaxHeap(); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input   []int
		wantMax int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 1},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 6},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		got, err := p.Max()
		if err != nil {
			t.Error(err)
		}
		if got != tt.wantMax {
			t.Errorf("%v. got %v, want %v", i, got, tt.wantMax)
		}
	}
}

func TestDelMaxAndInsert(t *testing.T) {
	tests := []struct {
		input       []int
		wantDelMax  int
		insertValue int
		wantMax     int
	}{
		{[]int{0}, 0, 1, 1},
		{[]int{0, 1}, 1, -1, 0},
		{[]int{0, 1, 2, 3, 4, 5}, 5, 10, 10},
		{[]int{0, 1, 2, 3, 4, 5}, 5, 3, 4},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		gotDelMax, err := p.DelMax()
		if err != nil {
			t.Error(err)
		}
		if gotDelMax != tt.wantDelMax {
			t.Errorf("%v. got %v, want %v", i, gotDelMax, tt.wantDelMax)
		}

		p.Insert(tt.insertValue)
		if isMaxHeap := p.isMaxHeap(); !isMaxHeap {
			t.Errorf("%v. got %v, want %v", i, isMaxHeap, true)
		}
		gotMax, err := p.Max()
		if err != nil {
			t.Error(err)
		}
		if gotMax != tt.wantMax {
			t.Errorf("%v. got %v, want %v", i, gotMax, tt.wantMax)
		}
	}
}
