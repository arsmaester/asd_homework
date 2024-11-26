package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 18, 20, 22}
	targets := []int{5, 2, 12, 7}

	var wg sync.WaitGroup
	results := make(chan int, len(targets))

	for _, target := range targets {
		wg.Add(1)
		go BinarySearch(arr, target, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		if result == -1 {
			fmt.Println("Target not found")
		} else {
			fmt.Printf("Target found at index: %dn", result)
		}
	}
}

func BinarySearch(arr []int, target int, wg *sync.WaitGroup, results chan<- int) int {
	defer wg.Done()
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
