package linearProbingHash

import (
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		input map[string]int
		key   string
		want  bool
	}{
		//{map[string]int{"a":0}, "a", true},
		//{map[string]int{"a":0}, "b", false},
		{map[string]int{"abc": 3, "bced": 4, "cedff": 5}, "bced", true},
	}
	for i, tt := range tests {
		s := New(len(tt.input))
		for k, v := range tt.input {
			s.Put(&k, &v)
		}
		if got := s.Contains(&tt.key); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestGEtt(t *testing.T) {
	s := New(3)
	k1 := "abc"
	k2 := "bced"
	k3 := "cedff"
	v1 := 3
	v2 := 4
	v3 := 5
	s.Put(&k1, &v1)
	s.Put(&k2, &v2)
	s.Put(&k3, &v3)

	query1 := "abc"
	query2 := "bced"
	query3 := "cedff"
	t.Log(s.Contains(&query1))
	t.Log(s.Contains(&query2))
	t.Log(s.Contains(&query3))
}

func TestPointer(t *testing.T) {
	a := "aaa"
	b := "aaa"
	pa := &a
	pb := &b
	t.Log(a == b)
	t.Log(pa == pb)
}

//func TestGet(t *testing.T) {
//	tests := []struct {
//		input map[int]int
//		key   int
//		want  int
//	}{
//		{map[int]int{0: 0}, 0, 0},
//		{map[int]int{0: 0, 1: 1, 2: 2}, 1, 1},
//		{map[int]int{0: 0, 1: 1, 2: 2}, 3, 0},
//	}
//	for i, tt := range tests {
//		s := New()
//		for k, v := range tt.input {
//			s.Put(k, v)
//		}
//
//		got, _ := s.Get(tt.key)
//
//		if got != tt.want {
//			t.Errorf("%v. got %v, want %v", i, got, tt.want)
//		}
//	}
//}
//
//func TestPut(t *testing.T) {
//	tests := []struct {
//		input map[int]int
//		key   int
//		value int
//		want  int
//	}{
//		{map[int]int{0: 0}, 0, 1, 1},
//		{map[int]int{0: 0}, 1, 1, 1},
//		{map[int]int{0: 0, 1: 1, 3: 3}, 2, 2, 2},
//	}
//	for i, tt := range tests {
//		s := New()
//		for k, v := range tt.input {
//			s.Put(k, v)
//		}
//
//		s.Put(tt.key, tt.value)
//
//		got, err := s.Get(tt.key)
//		if err != nil {
//			t.Error(err)
//		}
//		if got != tt.want {
//			t.Errorf("%v. got %v, want %v", i, got, tt.want)
//		}
//	}
//}
//
//func TestDelete(t *testing.T) {
//	tests := []struct {
//		input     map[int]int
//		deleteKey int
//		getKey    int
//		want      int
//	}{
//		{map[int]int{2: 2}, 2, 2, 0},
//		{map[int]int{1: 1, 2: 2}, 1, 2, 2},
//		{map[int]int{1: 1, 2: 2}, 1, 1, 0},
//		{map[int]int{1: 5, 2: 1, 3: 7, 4: 0, 5: 2, 6: 6, 7: 8}, 2, 7, 8},
//	}
//	for i, tt := range tests {
//		s := New()
//		for k, v := range tt.input {
//			s.Put(k, v)
//		}
//
//		s.Delete(tt.deleteKey)
//
//		got, _ := s.Get(tt.getKey)
//		if got != tt.want {
//			t.Errorf("%v. got %v, want %v", i, got, tt.want)
//		}
//	}
//}
