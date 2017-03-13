package sort

/**
this is a simple bubble sort
*/

func BubbleSort(arr []int) {
	if nil == arr || len(arr) == 0 {
		return
	}

	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				Swap(&arr, i, j)
			}
		}
	}
}
