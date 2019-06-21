package utils

import (
	"math"
)

// Swap exchanges s[i] and s[j].
// 交换切片中的两个元素。
// Assume:
// s != nil
// i >= 0 && i < len(s)
// j >= 0 && j < len(s)
func Swap(s []interface{}, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// IsSameSlice determines two slice is it the same.
// 判断两个切片是否相同。
func IsSameSlice(a, b []interface{}) bool {
	if len(a) != len(b) {
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
// 两个 2x2 整数矩阵相乘，返回一个新矩阵。
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
// 求整数切片的平均值。
func Average(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum / len(s)
}

// CopySlice returns a new slice copy of s.
// 复制一个新切片返回。
func CopySlice(s []interface{}) []interface{} {
	if s == nil {
		return nil
	}

	newSlice := make([]interface{}, cap(s))
	copy(newSlice, s)
	return newSlice
}

// FindMax returns max number from s.
// 返回整数切片的最大值。
func FindMax(s []int) int {
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

// GCD returns greatest common divisor with p and q.
// 返回 p q 的最大公因数。
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
// 原地反转切片。
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
// 判读是否是素数。
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
// 用牛顿迭代法求开方。
// Assume:
// c >= 0
func Sqrt(c float64) float64 {
	err := 1e-15
	t := c
	for math.Abs(t-c/t) > err*t {
		t = (c/t + t) / 2.0
	}
	return t
}

// Hypotenuse calculate the hypotenuse of a triangle.
// 求三角形斜边的值。
func Hypotenuse(a, b float64) float64 {
	return Sqrt(a*a + b*b)
}

// Harmonic calculate the Harmonic number.
// 计算调和级数。
func Harmonic(n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}
	return sum
}
