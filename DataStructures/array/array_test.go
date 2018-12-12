package array

import (
	"testing"

	"GoAlgorithms/utils"
)

func TestGet(t *testing.T) {
	tests := []struct {
		values []interface{}
	}{
		{[]interface{}{0, 1}},
	}
	for i, tt := range tests {
		a := New(cap(tt.values))
		for _, v := range tt.values {
			a.Append(v)
		}

		for k, v := range tt.values {
			got, err := a.Get(k)
			if err != nil {
				t.Error(err)
			}

			if got != v {
				t.Errorf("%v. got %v, want %v", i, got, v)
			}
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		values   []interface{}
		index    int
		setValue interface{}
	}{
		{[]interface{}{0}, 0, 1},
		{[]interface{}{0, 1, 2, 3}, 3, 4},
	}
	for i, tt := range tests {
		a := New(cap(tt.values))
		for _, v := range tt.values {
			a.Append(v)
		}

		if err := a.Set(tt.index, tt.setValue); err != nil {
			t.Error(err)
		}

		got, err := a.Get(tt.index)
		if err != nil {
			t.Error(err)
		}

		if got != tt.setValue {
			t.Errorf("%v. got %v, want %v", i, got, tt.setValue)
		}
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		cap    int
		values []interface{}
	}{
		{4, []interface{}{0, 1, 2, 3}},
		{4, []interface{}{'a', 'b', 'c', 'd'}},
		{2, []interface{}{0, 1, 2, 3}},
	}
	for i, tt := range tests {
		a := New(tt.cap)

		for _, v := range tt.values {
			a.Append(v)
		}

		if !utils.IsSameSlice(a.elements, tt.values) {
			t.Errorf("%v. got %v, want %v", i, a.elements, tt.values)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		cap  int
		give []interface{}
		want []interface{}
		i    int
		v    interface{}
	}{
		{4, []interface{}{0, 1, 3}, []interface{}{0, 1, 2, 3}, 2, 2},
		{4, []interface{}{'a', 'b', 'd'}, []interface{}{'a', 'b', 'c', 'd'}, 2, 'c'},
		{2, []interface{}{0, 2}, []interface{}{0, 1, 2, nil}, 1, 1},
	}
	for i, tt := range tests {
		a := New(tt.cap)
		for _, v := range tt.give {
			a.Append(v)
		}

		if err := a.Insert(tt.i, tt.v); err != nil {
			t.Error(err)
		}

		if !utils.IsSameSlice(a.elements, tt.want) {
			t.Errorf("%v. got %v, want %v", i, a.elements,
				tt.want)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		cap  int
		give []interface{}
		want []interface{}
		i    int
	}{
		{4, []interface{}{0, 1, 2, 3}, []interface{}{0, 1, 3, nil}, 2},
		{4, []interface{}{'a', 'b', 'c', 'd'}, []interface{}{'a', 'b', 'd', nil}, 2},
		{8, []interface{}{0, 1, nil, nil, nil, nil, nil, nil}, []interface{}{0, nil, nil, nil}, 1},
	}
	for i, tt := range tests {
		a := New(cap(tt.give))
		for k, v := range tt.give {
			a.elements[k] = v
			if v != nil {
				a.len++
			}
		}

		if err := a.Delete(tt.i); err != nil {
			t.Error(err)
		}

		if !utils.IsSameSlice(a.elements, tt.want) {
			t.Errorf("%v. got %v, want %v", i, a.elements,
				tt.want)
		}
	}
}

func TestResize(t *testing.T) {
	tests := []struct {
		cap  int
		give []interface{}
		want []interface{}
	}{
		{4 * 2, []interface{}{0, 1, 2, 3}, []interface{}{0, 1, 2, 3, nil, nil, nil, nil}},
		{8 / 2, []interface{}{0, 1, nil, nil, nil, nil, nil, nil}, []interface{}{0, 1, nil, nil}},
	}
	for i, tt := range tests {
		a := New(cap(tt.give))
		for k, v := range tt.give {
			a.elements[k] = v
			if v != nil {
				a.len++
			}
		}

		a.resize(tt.cap)

		if !utils.IsSameSlice(a.elements, tt.want) {
			t.Errorf("%v. got %v, want %v", i, a.elements, tt.want)
		}
	}
}
