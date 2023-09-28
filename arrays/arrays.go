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

func ArePermutations(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	lettersCounter := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		lettersCounter[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		lettersCounter[s2[i]]--
		if lettersCounter[s2[i]] < 0 {
			return false
		}
	}
	return true
}

func urlify(s []byte, trueLength int) {
	j := len(s) - 1
	i := trueLength - 1
	for ; i >= 0 && i < j; i-- {
		if s[i] == ' ' {
			s[j] = '0'
			s[j-1] = '2'
			s[j-2] = '%'
			j -= 3
		} else {
			s[j], j = s[i], j-1
		}
	}
}

func isAlphaNumeric(b byte) bool {
	if (b < 'a' || b > 'z') && (b < 'A' || b > 'Z') && (b < '0' || b > '9') {
		return false
	}
	return true
}

func isPermutationOfPalindrome(s []byte) bool {
	s = []byte(strings.ToLower(string(s)))
	countOfOdds := 0
	letterCountsHash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letter := s[i]
		if !isAlphaNumeric(letter) {
			continue
		}

		letterCountsHash[letter] += 1
		if letterCountsHash[letter]%2 == 0 {
			countOfOdds -= 1
		} else {
			countOfOdds += 1
		}
	}

	return (countOfOdds == 1)
}
