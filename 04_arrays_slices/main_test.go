package arraysslices

import "testing"

func TestSum(t *testing.T) {

	t.Run("Run a dynamic collection of 3 numbers", func(t *testing.T) {
		a := []int{1, 2, 3}

		sum := Sum(a)
		expected := 6

		if sum != expected {
			t.Errorf("Expected %d expected but got %d", expected, sum)
		}
	})
}
