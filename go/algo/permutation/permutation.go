package permutation

import "fmt"

// Simple permutation of all elements of the array
func perm(arr []int, k int, control *[]int, res *[]int) {
	if k == len(arr) {
		fmt.Println(res)
	} else {
		for i := 0; i < len(arr); i++ {
			if (*control)[i] == 0 {
				(*control)[i] = 1
				(*res)[k] = arr[i]
				perm(arr, k+1, control, res)
				(*control)[i] = 0
			}
		}
	}
}
