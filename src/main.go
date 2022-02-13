package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	arr := generateRandomSlice(11)
	arr2 := arr
	fmt.Printf("input:%v\n\n", arr)

	start := time.Now()
	bubble_sort(arr)
	fmt.Printf("sorted(loop):     %v\n", arr)
	elapsed := time.Since(start)
	log.Printf("It took %s", elapsed)
	fmt.Println("")

	start = time.Now()
	bubble_sort_recursion(arr2)
	fmt.Printf("sorted(recursion):%v\n", arr2)
	elapsed = time.Since(start)
	log.Printf("It took %s", elapsed)
}

func bubble_sort(arr []int) {
	for round := 1; round <= len(arr); round++ {
		end := len(arr) - round
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
			// fmt.Printf("round%d:%v\n", round, arr)
		}
	}
}

func bubble_sort_recursion(arr []int) {
	round := 1
	bubble_sort_recursion_help1(arr, round)
}

func bubble_sort_recursion_help1(arr []int, round int) {
	if round > len(arr) {
		return
	}

	bubble_sort_recursion_help2(arr, 0, round)

	bubble_sort_recursion_help1(arr, round+1)
}

func bubble_sort_recursion_help2(arr []int, i int, round int) {
	if i >= len(arr)-round {
		return
	}

	if arr[i] > arr[i+1] {
		swap(arr, i, i+1)
	}

	bubble_sort_recursion_help2(arr, i+1, round)
}

func swap(arr []int, left int, right int) {
	arr[left] = arr[left] ^ arr[right]
	arr[right] = arr[left] ^ arr[right]
	arr[left] = arr[left] ^ arr[right]
}

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}
