package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})
}

func TesSumAll(t *testing.T) {
	firstGroup := []int{1, 2}
	secondGroup := []int{3, 4}
	thirdGroup := []int{5}

	got := SumAll(firstGroup, secondGroup, thirdGroup)
	want := []int{3, 7, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
