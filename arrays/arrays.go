package arrays

import (
	"strconv"
	"strings"
)

func ArrayToString(array []int) string {
	stringArray := []string{}
	for i := range array {
		number := array[i]
		text := strconv.Itoa(number)
		stringArray = append(stringArray, text)
	}
	return "[" + strings.Join(stringArray, " ") + "]"
}

func MergeSort(array *[]int, indent int) {
	arr := *array
	if len(arr) > 1 {
		mid := len(arr) / 2

		L := make([]int, mid)
		copy(L[:], arr[:mid])
		R := make([]int, len(arr)-mid)
		copy(R[:], arr[mid:])
		MergeSort(&L, indent+1)
		MergeSort(&R, indent+1)

		i, j, k := 0, 0, 0
		for i < len(L) && j < len(R) {
			if L[i] <= R[j] {
				arr[k] = L[i]
				i++
			} else {
				arr[k] = R[j]
				j++
			}
			k++
		}
		for i < len(L) {
			arr[k] = L[i]
			i++
			k++
		}
		for j < len(R) {
			arr[k] = R[j]
			j++
			k++
		}
	}
}

func partition(arr []int, low int, high int) int {
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

func QuickSort(arr []int, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		QuickSort(arr, low, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, high)
	}
}
