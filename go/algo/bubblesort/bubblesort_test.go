package bubblesort

import "testing"

func check(array []int) bool {
	for i := 0; i < len(array); i++ {
		if i != len(array)-1 && array[i] > array[i+1] {
			return false
		}
	}

	return true
}

func TestSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Empty array",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Sorted array",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Reverse sorted array",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Unsorted array",
			input: []int{3, 1, 4, 2, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.input)
			if !check(tt.input) {
				t.Errorf("Sort() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
