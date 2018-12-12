package doublyCycleLinkedList

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		nodes []*node
		want  string
	}{
		{[]*node{NewNode(0)}, "0"},
		{[]*node{NewNode(2), NewNode(1), NewNode(0)}, "0->1->2"},
	}
	for i, tt := range tests {
		l := New()
		for _, v := range tt.nodes {
			if err := l.Insert(v, l.head); err != nil {
				t.Error(err)
			}
		}

		if got := l.Show(); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		nodes []*node
		index int
		want  string
	}{
		{[]*node{NewNode(0)}, 1, ""},
		{[]*node{NewNode(2), NewNode(1), NewNode(0)}, 1, "1->2"},
		{[]*node{NewNode(2), NewNode(1), NewNode(0)}, 3, "0->1"},
	}
	for i, tt := range tests {
		l := New()
		for _, v := range tt.nodes {
			if err := l.Insert(v, l.head); err != nil {
				t.Error(err)
			}
		}

		node := l.head
		for index := 0; index < tt.index; index++ {
			node = node.next
		}

		if err := l.Delete(node); err != nil {
			t.Error(err)
		}

		if got := l.Show(); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func TestShow(t *testing.T) {
	tests := []struct {
		nodes []*node
		want  string
	}{
		{[]*node{NewNode(0)}, "0"},
		{[]*node{NewNode(2), NewNode(1), NewNode(0)}, "0->1->2"},
	}
	for i, tt := range tests {
		l := New()
		for _, v := range tt.nodes {
			if err := l.Insert(v, l.head); err != nil {
				t.Error(err)
			}
		}

		if got := l.Show(); got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
