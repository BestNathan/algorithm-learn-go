package main

func SelectionSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		selectedIndex := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[selectedIndex] {
				selectedIndex = j
			}
		}

		arr[i], arr[selectedIndex] = arr[selectedIndex], arr[i]
	}

	return arr
}