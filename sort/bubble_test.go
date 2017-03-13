package sort

import (
	_ "./"
	"testing"
)

var intArr = []int{1, 3, 6, 2, 7, 4}

func TestBubbleSort(t *testing.T) {
	BubbleSort(intArr)
}
