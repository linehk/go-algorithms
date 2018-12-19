package maxPriorityQueue

import (
	"testing"
)

func TestNewByArray(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{}},
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
		input []int
		want  int
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
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestDelMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 1},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 6},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		got, err := p.DelMax()
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
		if isMaxHeap := p.isMaxHeap(); !isMaxHeap {
			t.Errorf("%v. got %v, want %v", i, isMaxHeap, true)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		input       []int
		insertValue int
	}{
		{[]int{0}, 0},
		{[]int{0, 0}, 1},
		{[]int{0, 1, 2, 3}, 4},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		if err := p.Insert(tt.insertValue); err != nil {
			t.Error(err)
		}
		if got := p.isMaxHeap(); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}
