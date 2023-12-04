package arrays

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
