package sequentialSearchST

import (
	"testing"
)

func TestSequentialSearchST(t *testing.T) {
	st := New()
	keys := []int{0, 1}
	values := []int{1, 2}
	st.Put(keys[0], values[0])
	v, err := st.Get(keys[0])
	if err != nil {
		t.Error(err)
	}
	if v != values[0] {
		t.Errorf("got %v, want %v", v, values[0])
	}

	isContains, err := st.Contains(keys[0])
	if !isContains {
		t.Errorf("got %v, want %v", !isContains, true)
	}

	st.Put(keys[1], values[1])
	v, err = st.Get(keys[1])
	if err != nil {
		t.Error(err)
	}
	if v != values[1] {
		t.Errorf("got %v, want %v", v, values[1])
	}

	st.Delete(keys[1])
	v, err = st.Get(keys[1])
	if err == nil {
		t.Errorf("key should be delete")
	}
}
