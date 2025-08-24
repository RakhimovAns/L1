package main

import "fmt"

func BinarySearch(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}

	l, r := 0, len(arr)-1

	for l+1 < r {
		mid := (l + r) / 2

		if arr[mid] < k {
			l = mid
		} else {
			r = mid
		}
	}

	if arr[r] == k {
		return r
	}
	return -1
}

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	//var k = 10

	fmt.Println(BinarySearch(arr, 0))
	fmt.Println(BinarySearch(arr, 3))
	fmt.Println(BinarySearch(arr, 10))
	fmt.Println(BinarySearch([]int{}, 1))
	fmt.Println(BinarySearch([]int{1}, 1))
	fmt.Println(BinarySearch([]int{1}, 2))
	fmt.Println(BinarySearch([]int{1, 3, 5}, 2))
}
