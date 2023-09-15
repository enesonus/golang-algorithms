package main

import (
	"fmt"
)

func main() {
	arr := []int{3,8,5,4,1,9,-2,0,4,55,5,5,9}
	
	quick_sort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low-1
	j := low
	for ;j < high; j++{
		if arr[j] < pivot{
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i+1
}

func quick_sort(arr []int, low, high int) {

	if len(arr) <= 1 || low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quick_sort(arr, low, pivotIndex-1)
	quick_sort(arr, pivotIndex+1, high)

}