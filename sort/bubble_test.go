package sort

import (
	"fmt"
	"testing"
)

var intArray = []int{1, 8, 3, 6, 2, 7, 4}

func TestBubbleSort(t *testing.T) {
	fmt.Println("before sort : ", intArray)
	BubbleSort(intArray)
	fmt.Println("after sort : ", intArray)
}
