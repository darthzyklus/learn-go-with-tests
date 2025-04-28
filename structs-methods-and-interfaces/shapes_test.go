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

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("area in rectangle", func(t *testing.T) {
		rect := Rectangle{
			width:  12.0,
			height: 6.0,
		}

		want := 72.0

		checkArea(t, rect, want)
	})

	t.Run("area in circle", func(t *testing.T) {
		circle := Circle{
			Radius: 10.0,
		}

		want := math.Pi * 100

		checkArea(t, circle, want)

	})

}
