package separateChainingHash

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		input map[string]int
		key   string
		want  bool
	}{
		{map[string]int{"a": 0}, "a", true},
		{map[string]int{"a": 0}, "b", false},
		{map[string]int{"abc": 3, "bced": 4, "cedff": 5}, "bced", true},
	}
	for i, tt := range tests {
		s := New(len(tt.input))
		for k, v := range tt.input {
			s.Put(k, v)
		}
		if got := s.Contains(tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		input map[string]int
		key   string
		want  int
	}{
		{map[string]int{"a": 0}, "a", 0},
		{map[string]int{"a": 0, "b": 1, "c": 2}, "b", 1},
		{map[string]int{"a": 0, "b": 1, "c": 2}, "d", 0},
	}
	for i, tt := range tests {
		s := New(len(tt.input))
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
		input map[string]int
		key   string
		value int
		want  int
	}{
		{map[string]int{"a": 0}, "b", 1, 1},
		{map[string]int{"a": 0, "b": 1, "c": 2}, "c", 2, 2},
	}
	for i, tt := range tests {
		s := New(len(tt.input))
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
		input     map[string]int
		deleteKey string
		getKey    string
		want      int
	}{
		{map[string]int{"a": 0}, "a", "a", 0},
		{map[string]int{"a": 0, "b": 1}, "a", "b", 1},
		{map[string]int{"a": 0, "b": 1}, "a", "a", 0},
		{map[string]int{"a": 0, "b": 1, "c": 2}, "b", "c", 2},
	}
	for i, tt := range tests {
		s := New(len(tt.input))
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

func TestResize(t *testing.T) {
	tests := []struct {
		input map[string]int
	}{
		{map[string]int{"a": 0}},
		{map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4}},
	}
	for i, tt := range tests {
		l := len(tt.input)
		s := New(l)
		for k, v := range tt.input {
			s.Put(k, v)
		}

		s.resize(l * 2)

		for k, v := range tt.input {
			got, err := s.Get(k)
			if err != nil {
				t.Error(err)
			}
			if got != v {
				t.Errorf("%v. got %v, want %v", i, got, v)
			}
		}
	}
}
