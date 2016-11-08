package main

import (
	"fmt"
	"testing"
)

func TestMergeSortNewArray(t *testing.T) {
	array := []int{1, 3, 5, 2, 4, 6}
	sorted, inversions := MergeSortNewArray(array)
	fmt.Println("The number of inversions:", inversions)
	if !testSliceEquality([]int{1, 2, 3, 4, 5, 6}, sorted) {
		t.Error("Expected [1 2 3 4 5 6] but got:", sorted)
	}

	arr := []int{1, 3, 5, 6, 2, 4, 7}
	s, i := MergeSortNewArray(arr)
	fmt.Println("The number of inversions pt2:", i)
	if !testSliceEquality([]int{1, 2, 3, 4, 5, 6, 7}, s) {
		t.Error("Expected [1 2 3 4 5 6 7] but got:", s)
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
