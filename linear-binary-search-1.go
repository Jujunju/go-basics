package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	data := []int{12, 22, 25, 34, 64}
	fmt.Println("Linear Search posisi 34:", linearSearch(data, 34))
	fmt.Println("Binary Search posisi 34:", binarySearch(data, 34))
}
