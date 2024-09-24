package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %d expected but got %d", want, got)
		}
	}

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
		checkSum(t, sum, expected)
	})

	t.Run("Add the slices passed but do not include the first element", func(t *testing.T) {
		sum := SumAllTrail([]int{1, 2, 3}, []int{3, 5, 6, 8})
		expected := []int{5, 19}
		checkSum(t, sum, expected)
	})

	t.Run("Safely sum slices with no elements", func(t *testing.T) {
		sum := SumAllTrail([]int{1, 2, 3}, []int{})
		expected := []int{5, 0}
		checkSum(t, sum, expected)
	})
}
