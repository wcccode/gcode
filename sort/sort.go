package sort

func Swap(arr *[]int, src, dest int) {
	temp := (*arr)[src]
	(*arr)[src] = (*arr)[dest]
	(*arr)[dest] = temp
}
