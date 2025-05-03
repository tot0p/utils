package tree

import "testing"

func TestBinaryTree(t *testing.T) {
	// Create a new tree and insert some values
	root := NewNode(10)
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)

	// Test searching for existing values
	tests := []struct {
		value    int
		expected bool
	}{
		{10, true},
		{5, true},
		{15, true},
		{3, true},
		{7, true},
		{20, false},
		{2, false},
	}

	for _, tt := range tests {
		if got := root.Search(tt.value); got != tt.expected {
			t.Errorf("Search(%d) = %v, want %v", tt.value, got, tt.expected)
		}
	}
}
