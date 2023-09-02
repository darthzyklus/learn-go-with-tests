package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("2 + 2 should return 4", func(t *testing.T) {
		sum := Add(2, 2)

		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
	t.Run("1 + 5 should return 6", func(t *testing.T) {
		sum := Add(1, 5)

		expected := 6

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
