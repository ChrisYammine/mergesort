package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func MergeSort(array []int) {
	l := len(array)
	temp := make([]int, l) // Rather than reallocate a new array for each merge, we'll just use this one
	mergeSort(array, temp, 0, l-1)
}

func mergeSort(array, temp []int, left, right int) {
	if left >= right {
		return
	} // finished case
	middle := (left + right) / 2
	mergeSort(array, temp, left, middle)
	mergeSort(array, temp, middle+1, right)
	merge(array, temp, left, right)
}

func merge(array, temp []int, leftStart, rightEnd int) {
	leftEnd := (rightEnd + leftStart) / 2
	rightStart := leftEnd + 1
	size := rightEnd - leftStart + 1

	left := leftStart
	right := rightStart
	index := leftStart

	for left <= leftEnd && right <= rightEnd {
		if array[left] <= array[right] {
			temp[index] = array[left]
			left++
		} else {
			temp[index] = array[right]
			right++
		}
		index++
	}

	// copy over remaining elements from left (only one side will have remaining elements)
	for left <= leftEnd {
		temp[index] = array[left]
		left++
		index++
	}
	// copy over remaining elements from right (only one side will have remaining elements)
	for right <= rightEnd {
		temp[index] = array[right]
		right++
		index++
	}

	// copy our changes into original array
	for i := leftStart; i < size; i++ {
		array[i] = temp[i]
	}
	fmt.Println("Original array:", array)
}

// Returning new array
func MergeSortNewArray(array []int, inv int) ([]int, int) {
	if len(array) < 2 {
		return array, 0
	}

	middle := len(array) / 2
	a, inv1 := MergeSortNewArray(array[:middle], inv)
	b, inv2 := MergeSortNewArray(array[middle:], inv)
	return merge1(a, b, inv1+inv2)
}

func merge1(a, b []int, inversions int) ([]int, int) {
	result := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result[i+j] = a[i]
			i++
		} else {
			result[i+j] = b[j]
			j++
		}
	}

	inv := inversions + i + j

	for i < len(a) {
		result[i+j] = a[i]
		i++
	}
	for j < len(b) {
		result[i+j] = b[j]
		j++
	}

	return result, inv
}

func main() {
	fileNameScanner := bufio.NewScanner(os.Stdin)
	fileNameScanner.Scan()
	file, err := os.Open("./" + fileNameScanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	datasets := make([][]int, n)
	for i, _ := range datasets {
		scanner.Scan() // throw away the number of elements in the array since I can't properly scan using it on hackerrank
		scanner.Scan() // the space delimited array of numbers
		array := strings.Split(scanner.Text(), " ")
		numberArray := make([]int, len(array))
		for i, v := range array {
			num, _ := strconv.Atoi(v)
			numberArray[i] = num
		}
		datasets[i] = numberArray
	}
	for _, v := range datasets {
		_, inversions := MergeSortNewArray(v, 0)
		fmt.Println(inversions)
	}
}
