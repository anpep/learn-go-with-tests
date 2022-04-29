package structs_methods_and_interfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		rectangle         Rectangle
		expectedPerimeter float64
	}{
		{rectangle: Rectangle{Width: 0, Height: 0}, expectedPerimeter: 0},
		{rectangle: Rectangle{Width: 10, Height: 10}, expectedPerimeter: 40},
	}

	for _, test := range perimeterTests {
		calculatedPerimeter := test.rectangle.Perimeter()
		if calculatedPerimeter != test.expectedPerimeter {
			t.Errorf("expected perimeter of %#v to be %g but got %g",
				test.rectangle, test.expectedPerimeter, calculatedPerimeter)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape        Shape
		expectedArea float64
	}{
		{shape: Rectangle{Width: 0, Height: 10}, expectedArea: 0},
		{shape: Rectangle{Width: 10, Height: 10}, expectedArea: 100},
		{shape: Circle{Radius: 0}, expectedArea: 0},
		{shape: Circle{Radius: 10}, expectedArea: 100 * math.Pi},
		{shape: Triangle{Base: 0, Height: 0}, expectedArea: 0},
		{shape: Triangle{Base: 12, Height: 6}, expectedArea: 36},
	}

	for _, test := range areaTests {
		calculatedArea := test.shape.Area()
		if calculatedArea != test.expectedArea {
			t.Errorf("expected area of %#v to be %g but got %g",
				test.shape, test.expectedArea, calculatedArea)
		}
	}
}
