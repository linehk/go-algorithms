package redBlackBST

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
		{map[int]int{0: 0}, 1, 1, 1},
		{map[int]int{0: 0, 1: 1, 3: 3}, 2, 2, 2},
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
		{map[int]int{1: 5, 2: 1, 3: 7, 4: 0, 5: 2, 6: 6, 7: 8}, 2, 7, 8},
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

func TestDeleteMinAndMax(t *testing.T) {
	tests := []struct {
		input   map[int]int
		wantMin int
		wantMax int
	}{
		{map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 1, 2},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		if err := s.DeleteMin(); err != nil {
			t.Error(err)
		}
		gotMin, err := s.Min()
		if err != nil {
			t.Error(err)
		}
		if gotMin != tt.wantMin {
			t.Errorf("%v. got %v, want %v", i, gotMin, tt.wantMin)
		}

		if err := s.DeleteMax(); err != nil {
			t.Error(err)
		}
		gotMax, err := s.Max()
		if err != nil {
			t.Error(err)
		}
		if gotMax != tt.wantMax {
			t.Errorf("%v. got %v, want %v", i, gotMax, tt.wantMax)
		}
	}
}

func TestFloorAndCeiling(t *testing.T) {
	tests := []struct {
		input       map[int]int
		key         int
		wantFloor   int
		wantCeiling int
	}{
		{map[int]int{0: 0, 1: 1, 3: 3, 4: 4}, 2, 1, 3},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}

		gotFloor, err := s.Floor(tt.key)
		if err != nil {
			t.Error(err)
		}
		if gotFloor != tt.wantFloor {
			t.Errorf("%v. got %v, want %v", i, gotFloor, tt.wantFloor)
		}

		gotCeiling, err := s.Ceiling(tt.key)
		if err != nil {
			t.Error(err)
		}
		if gotCeiling != tt.wantCeiling {
			t.Errorf("%v. got %v, want %v", i, gotFloor, tt.wantFloor)
		}
	}
}

func TestSelect(t *testing.T) {
	tests := []struct {
		input map[int]int
		k     int
		want  int
	}{
		{map[int]int{0: 0, 1: 1, 2: 2}, 1, 1},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		got, err := s.Select(tt.k)
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestSizeByRange(t *testing.T) {
	tests := []struct {
		input    map[int]int
		lo       int
		hi       int
		wantSize int
	}{
		{map[int]int{0: 0, 1: 1, 2: 2, 3: 3}, 0, 3, 4},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}
		if gotSize := s.SizeByRange(tt.lo, tt.hi); gotSize != tt.wantSize {
			t.Errorf("%v. got %v, want %v", i, gotSize, tt.wantSize)
		}
	}
}

func TestHeight(t *testing.T) {
	tests := []struct {
		input map[int]int
		want  int
	}{
		{map[int]int{}, -1},
		{map[int]int{0: 0}, 0},
		{map[int]int{0: 0, 1: 1}, 1},
		{map[int]int{0: 0, 1: 1, 2: 2}, 1},
	}
	for i, tt := range tests {
		s := New()
		for k, v := range tt.input {
			s.Put(k, v)
		}

		if got := s.Height(); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
