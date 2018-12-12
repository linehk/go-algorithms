package utils

import (
	"errors"
	"math"
)

// Swap exchanges s[i] and s[j].
func Swap(s []interface{}, i, j int) error {
	if s == nil {
		return errors.New("slice can't be nil")
	}
	if (i < 0 || i >= len(s)) || (j < 0 || j >= len(s)) {
		return errors.New("illegal index")
	}

	s[i], s[j] = s[j], s[i]

	return nil
}

// IsSameSlice determines two slice is it the same.
func IsSameSlice(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

// MatrixMultiply returns two matrix product.
func MatrixMultiply(a, b [][]int) [][]int {
	c := [][]int{
		{0, 0},
		{0, 0},
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			c[i][j] += a[i][j] * b[i][j]
		}
	}

	return c
}

// Average returns s average number.
func Average(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum / len(s)
}

// CopySlice returns a new slice copy of s.
func CopySlice(s []interface{}) []interface{} {
	if s == nil {
		return nil
	}

	newSlice := make([]interface{}, cap(s))
	for i, v := range s {
		newSlice[i] = v
	}
	return newSlice
}

// FindMax returns max number from s.
func FindMax(s []int) int {
	var max int
	for _, v := range s {
		max = v
		if v > max {
			max = v
		}
	}
	return max
}

// GCD returns greatest common divisor with p and q.
func GCD(p, q int) int {
	if p < 0 && q < 0 {
		return 0
	}
	if q == 0 {
		return p
	}

	r := p % q

	return GCD(q, r)
}

// Reverse reverses s.
func Reverse(s []interface{}) {
	if s == nil {
		return
	}

	l := len(s)
	for i := 0; i < l/2; i++ {
		Swap(s, i, l-i-1)
	}
}

// IsPrime determines n is it the prime.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Sqrt extracts a root by Newton's method.
func Sqrt(c float64) (float64, error) {
	if c < 0 {
		return 0, errors.New("argument can't be negative")
	}

	err := 1e-15
	t := c
	for math.Abs(t-c/t) > err*t {
		t = (c/t + t) / 2.0
	}
	return t, nil
}

// Hypotenuse calculate the hypotenuse of a triangle.
func Hypotenuse(a, b float64) (float64, error) {
	v, err := Sqrt(a*a + b*b)
	if err != nil {
		return 0, errors.New("argument can't negative")
	}
	return v, nil
}

// Harmonic calculate the Harmonic number.
func Harmonic(n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}
