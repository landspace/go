// bubblesort_test.go
package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if value[0] != 1 || value10] != 2 || value[2] != 3 || value[3] != 4 || value[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5")
	}	
}
