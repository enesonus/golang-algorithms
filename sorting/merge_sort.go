package main

func merge_sort(arr []int){
	if len(arr) <= 1{
		return}

	middle := len(arr)/2
	leftArr := make([]int, len(arr[:middle]))
	rightArr := make([]int, len(arr[middle:]))

	copy(leftArr, arr[:middle])
	copy(rightArr, arr[middle:])

	merge_sort(leftArr)
	merge_sort(rightArr)
	merge(leftArr, rightArr, arr)
	
}

func merge(leftArr, rightArr , arr []int){
	leftSize := len(leftArr)
	rightSize := len(rightArr)
	l,r := 0,0
	for l<leftSize && r<rightSize{
		if(leftArr[l] < rightArr[r]){
			arr[l+r] = leftArr[l]
			l++
		} else{
			arr[l+r] = rightArr[r]
			r++
		}
	}
	for l<leftSize{
		arr[l+r] = leftArr[l]
		l++
	}
	for r<rightSize{
		arr[l+r] = rightArr[r]
		r++
	}
}