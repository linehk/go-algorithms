package adjacencyMatrix

import (
	"testing"
)

func TestMultArray(t *testing.T) {
	m := make([][]int, 3)
	for i := range m {
		m[i] = make([]int, 4)
	}
	m[0][1] = 1
	t.Log(m)

	t.Log(m[0])
}
