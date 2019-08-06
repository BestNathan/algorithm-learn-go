package main

import (
	"fmt"
)

func main() {
	unsortedArr := []int{1,5,3,4,2,6,9,7,8,0}
	
	bubbleSortedArr := BubbleSort(unsortedArr)
	fmt.Printf("bubble sorted arr: %v\n", bubbleSortedArr)

	
}