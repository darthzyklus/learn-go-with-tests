package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type AreaTestCase struct {
	name  string
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {
	areaTests := []AreaTestCase{
		{name: "Rectangle", shape: Rectangle{width: 12.0, height: 6.0}, want: 72.0},
		{name: "Circle", shape: Circle{radius: 10.0}, want: math.Pi * 100},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
