package main

import (
	"log"
	"math/rand"
)

const MAX_UINT = ^uint(0) 
const MIN_UINT = 0 
const MAX_INT = int(MAX_UINT >> 1) 
const MIN_INT = -MAX_INT - 1

func main(){
	sortedArr := generateSortedIntSlice(1000, false)
	unSortedArr := make([]int, len(sortedArr))
	copy(unSortedArr, sortedArr)

	shuffle(unSortedArr)

	quick_sort(unSortedArr, 0, len(unSortedArr)-1)

	if isEqual(sortedArr, unSortedArr){
		log.Println("Quick Sort is OK!")
		shuffle(unSortedArr)
	} else {
		log.Println("Not sorted!")
	}

	merge_sort(unSortedArr)

	if isEqual(sortedArr, unSortedArr){
		log.Println("Merge Sort is OK!")
		shuffle(unSortedArr)
	} else {
		log.Println("Not sorted!")
	}
}

func shuffle(arr []int) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func isEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func generateSortedIntSlice(length int, isUnsigned bool) []int {
	var rangeMin int
	if isUnsigned {
		rangeMin = MIN_UINT
	}

	rangeMax := MAX_INT
	if rangeMin >= rangeMax {
		log.Fatal("Error: rangeMin should be less than rangeMax.")
		return nil
	}

	result := make([]int, length)
	var low, diff int = 0, 100
	for i := 0; i < length; i++ {
		// Generate a random number within the given range
		num := rand.Intn(diff) + low
		low = num
		result[i] = num
	}

	return result
}