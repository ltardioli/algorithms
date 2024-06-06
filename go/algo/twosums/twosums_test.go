package twosums

import "testing"

func TestTwoSumsSolution(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		array := []int{1, 5, 8, 3, 2}
		sum := 5
		expected1 := 3
		expected2 := 4
		i, j, _ := TwoSums(array, sum)
		if i != expected1 || j != expected2 {
			t.Fatalf("expected %d and % d, got %d and %d", expected1, expected2, i, j)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		array := []int{4, 5, 8, 3, 2}
		sum := 9
		expected1 := 0
		expected2 := 1
		i, j, _ := TwoSums(array, sum)
		if i != expected1 || j != expected2 {
			t.Fatalf("expected %d and % d, got %d and %d", expected1, expected2, i, j)
		}
	})
}
