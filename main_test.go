package main

import (
	"fmt"
	"testing"
)

func TestMergeSortNewArray(t *testing.T) {
	array := []int{5, 4, 3, 2, 1}
	sorted, inversions := MergeSortNewArray(array, 0)
	fmt.Println("The number of inversions:", inversions)
	if !testSliceEquality([]int{1, 2, 3, 4, 5}, sorted) {
		t.Error("Expected [1 2 3 4 5] but got:", sorted)
	}
}

// Helper to test slice equality
func testSliceEquality(first, second []int) bool {
	if first == nil && second == nil {
		return true
	}
	if first == nil || second == nil {
		return false
	}
	if len(first) != len(second) {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
