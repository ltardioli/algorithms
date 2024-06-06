package twosums

import "fmt"

// Find two numbers in an array that add up to a given target number
func TwoSums(arr []int, sum int) (int, int, error) {
	holder := make(map[int]int, len(arr))
	for i, value := range arr {
		holder[value] = i
	}

	for i, value := range arr {
		diff := sum - value
		value, ok := holder[diff]
		if ok {
			fmt.Printf("Elements at %d and %d has their sums as %d", i, value, sum)
			return i, holder[diff], nil
		}
	}

	return -1, -1, fmt.Errorf("no solution found")
}
