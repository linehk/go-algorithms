package minPriorityQueue

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
		if got := p.isMinHeap(); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 0},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 0},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		got, err := p.Min()
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestDelMin(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{0}, 0},
		{[]int{0, 1}, 0},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 0},
	}
	for i, tt := range tests {
		p := NewByArray(tt.input)
		got, err := p.DelMin()
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
		if isMinHeap := p.isMinHeap(); !isMinHeap {
			t.Errorf("%v. got %v, want %v", i, isMinHeap, true)
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
		p.Insert(tt.insertValue)
		if got := p.isMinHeap(); !got {
			t.Errorf("%v. got %v, want %v", i, got, true)
		}
	}
}
