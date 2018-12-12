package utils

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tests := []struct {
		give []interface{}
		i    int
		j    int
		want []interface{}
	}{
		{[]interface{}{1, 'a', "aa"}, 0, 2, []interface{}{"aa", 'a', 1}},
	}
	for i, tt := range tests {
		if err := Swap(tt.give, tt.i, tt.j); err != nil {
			t.Error(err)
		}

		if !IsSameSlice(tt.give, tt.want) {
			t.Errorf("%v. got %v, want %v", i, tt.give, tt.want)
		}
	}
}

func TestIsSameSlice(t *testing.T) {
	tests := []struct {
		s1   []interface{}
		s2   []interface{}
		want bool
	}{
		{[]interface{}{1, 2, 3}, []interface{}{1, 2, 3}, true},
	}
	for i, tt := range tests {
		if got := IsSameSlice(tt.s1, tt.s2); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestMatrixMultiply(t *testing.T) {
	tests := []struct {
		a    [][]int
		b    [][]int
		want [][]int
	}{
		{[][]int{{2, 2}, {2, 2}}, [][]int{{2, 2}, {2, 2}}, [][]int{{4, 4}, {4, 4}}},
	}
	for i, tt := range tests {
		got := MatrixMultiply(tt.a, tt.b)
		for k := 0; k < len(got); k++ {
			for j := 0; j < len(got); j++ {
				if got[k][j] != tt.want[k][j] {
					t.Errorf("%v. got %v, want %v",
						i, got, tt.want)
				}
			}
		}
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{[]int{2, 2, 2, 2}, 2},
	}
	for i, tt := range tests {
		if got := Average(tt.give); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{1, 2, 3}, []interface{}{1, 2, 3}},
	}
	for i, tt := range tests {
		got := CopySlice(tt.give)
		if !IsSameSlice(got, tt.want) {
			t.Errorf("%v. got %v, want %v",
				i, got, tt.want)
		}
	}
}

func TestFindMax(t *testing.T) {
	tests := []struct {
		give []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
	}
	for i, tt := range tests {
		if got := FindMax(tt.give); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		p    int
		q    int
		want int
	}{
		{4, 8, 4},
	}
	for i, tt := range tests {
		if got := GCD(tt.p, tt.q); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		give []interface{}
		want []interface{}
	}{
		{[]interface{}{1, 2, 3}, []interface{}{3, 2, 1}},
	}
	for i, tt := range tests {
		Reverse(tt.give)
		if !IsSameSlice(tt.give, tt.want) {
			t.Errorf("%v. got %v, want %v", i, tt.give, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		give int
		want bool
	}{
		{5, true},
	}
	for i, tt := range tests {
		if got := IsPrime(tt.give); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		give float64
		want float64
	}{
		{4.0, 2.0},
	}
	for i, tt := range tests {
		got, err := Sqrt(tt.give)
		if err != nil {
			t.Error(err)
		}

		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestHypotenuse(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{3.0, 4.0, 5.0},
	}
	for i, tt := range tests {
		got, err := Hypotenuse(tt.a, tt.b)
		if err != nil {
			t.Error(err)
		}
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestHarmonic(t *testing.T) {
	tests := []struct {
		give int
		want float64
	}{
		{3, 11.0 / 6.0},
	}
	for i, tt := range tests {
		if got := Harmonic(tt.give); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
