package structsmethodsinterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	}
	t.Run("Testing Perimeter of Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("Testing Perimeter of Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	shapes := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{10.0, 10.0}, 50.0},
	}

	for _, tt := range shapes {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("%#v Got %g want %g", tt.shape, got, tt.want)
		}
	}

}
