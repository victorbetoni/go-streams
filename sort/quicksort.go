package sort

func QuickSort[E any](arr []E, comparator func(E, E) int, low, high int) []E {
	if low < high {
		var p int
		arr, p = partition(arr, comparator, low, high)
		arr = QuickSort(arr, comparator, low, p-1)
		arr = QuickSort(arr, comparator, p+1, high)
	}
	return arr
}

func partition[E any](arr []E, comparator func(E, E) int, low, high int) ([]E, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if comparator(arr[j], pivot) == -1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
