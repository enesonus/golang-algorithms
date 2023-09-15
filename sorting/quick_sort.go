package main

// This is quicksort algorithm. You can find a good and brief
// explanation here: https://www.youtube.com/watch?v=COk73cpQbFQ

func quick_sort(arr []int, low, high int) {

	if len(arr) <= 1 || low >= high {
		return
	}

	pivotIndex := partition(arr, low, high)
	quick_sort(arr, low, pivotIndex-1)
	quick_sort(arr, pivotIndex+1, high)

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


func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
