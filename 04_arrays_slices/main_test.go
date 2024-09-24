package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Run a dynamic collection of 3 numbers", func(t *testing.T) {
		a := []int{1, 2, 3}

		sum := Sum(a)
		expected := 6

		if sum != expected {
			t.Errorf("Expected %d expected but got %d", expected, sum)
		}
	})

	t.Run("Add the slices passed", func(t *testing.T) {

		sum := SumAll([]int{1, 2, 3}, []int{3, 5})
		expected := []int{6, 8}

		if reflect.DeepEqual(sum, expected) {
			t.Errorf("Expected %v expected but got %v", expected, sum)
		}
	})

}
