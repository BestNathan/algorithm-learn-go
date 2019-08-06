package main

import (
	"fmt"
)

func main() {
	runSortAlgorithm()
}

func copySlice(arr []int) []int {
	newArr := make([]int, len(arr), cap(arr))
	copy(newArr, arr)
	return newArr
}

func runSortAlgorithm()  {
	unsortedArr := []int{1,5,3,4,2,6,9,7,8,0}
	
	// bubble sort
	fmt.Printf("before sorted arr: %v\n", unsortedArr)
	bubbleSortedArr := BubbleSort(copySlice(unsortedArr))
	fmt.Printf("bubble sorted arr: %v\n", bubbleSortedArr)

	// selection sort
	fmt.Printf("before sorted arr: %v\n", unsortedArr)
	seletionSortedArr := SelectionSort(copySlice(unsortedArr))
	fmt.Printf("selection sorted arr: %v\n", seletionSortedArr)
}