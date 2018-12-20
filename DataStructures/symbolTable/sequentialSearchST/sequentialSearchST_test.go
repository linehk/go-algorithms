package sequentialSearchST

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		input        map[int]int
		key          int
		wantContains bool
	}{
		{map[int]int{0: 0}, 0, true},
		{map[int]int{0: 0}, 1, false},
		{map[int]int{0: 0, 1: 1, 2: 2}, 2, true},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		if gotContains := s.Contains(tt.key); gotContains != tt.wantContains {
			t.Errorf("%v. got %v, want %v", i, gotContains, tt.wantContains)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		input map[int]int
		key   int
		want  int
	}{
		{map[int]int{0: 0}, 0, 0},
		{map[int]int{0: 0, 1: 1, 2: 2}, 1, 1},
		{map[int]int{0: 0, 1: 1, 2: 2}, 3, 0},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		got, _ := s.Get(tt.key)
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestPut(t *testing.T) {
	tests := []struct {
		input map[int]int
		key   int
		value int
		want  int
	}{
		{map[int]int{0: 0}, 0, 1, 1},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		s.Put(tt.key, tt.value)
		got, err := s.Get(tt.key)
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		input     map[int]int
		deleteKey int
		getKey    int
		want      int
	}{
		{map[int]int{2: 2}, 2, 2, 0},
		{map[int]int{1: 1, 2: 2}, 1, 2, 2},
		{map[int]int{1: 1, 2: 2}, 1, 1, 0},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		s.Delete(tt.deleteKey)
		got, _ := s.Get(tt.getKey)
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
