package main

import (
	"fmt"
	"sync"
)

// parallelBinarySearch performs a binary search and returns the index offset-adjusted.
func parallelBinarySearch(arr []int, target int, offset int, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			resultChan <- mid + offset
			return
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	resultChan <- -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	var wg sync.WaitGroup
	resultChan := make(chan int, 2)

	mid := len(arr) / 2

	wg.Add(2)
	go parallelBinarySearch(arr[:mid], target, 0, &wg, resultChan)
	go parallelBinarySearch(arr[mid:], target, mid, &wg, resultChan)

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		if result != -1 {
			fmt.Printf("Found target at index: %d\n", result)
			return
		}
	}
	fmt.Println("Target not found")
}
