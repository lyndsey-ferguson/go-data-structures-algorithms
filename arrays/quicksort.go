package arrays

import "cmp"

func partition[T cmp.Ordered](arr []T, low int, high int) int {
	pivot := arr[high]

	bestKnownCorrectPivotIndex := low - 1
	for index := low; index <= high-1; index++ {
		if arr[index] < pivot {
			bestKnownCorrectPivotIndex++
			arr[bestKnownCorrectPivotIndex], arr[index] = arr[index], arr[bestKnownCorrectPivotIndex]
		}
	}
	arr[bestKnownCorrectPivotIndex+1], arr[high] = arr[high], arr[bestKnownCorrectPivotIndex+1]

	return bestKnownCorrectPivotIndex + 1
}

func QuickSort[T cmp.Ordered](arr []T, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		QuickSort(arr, low, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, high)
	}
}
