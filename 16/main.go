package main

import "fmt"

func partition(arr []int, low, high int) int {
	pivot := arr[(low+high)/2]
	i, j := low-1, high+1
	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p)
		quickSort(arr, p+1, high)
	}
}

func quickSortStart(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func main() {
	fmt.Println(quickSortStart([]int{5, 6, 7, 2, 1, 0}))
}
