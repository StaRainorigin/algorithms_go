package solution

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int64
	}{
		{
			name:     "example 1",
			nums:     []int{1, 3, 1, 1, 2},
			expected: []int64{5, 0, 3, 4, 0},
		},
		{
			name:     "example 2",
			nums:     []int{0, 5, 3},
			expected: []int64{0, 0, 0},
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: []int64{0},
		},
		{
			name:     "two same",
			nums:     []int{1, 1},
			expected: []int64{1, 1},
		},
		{
			name:     "three same",
			nums:     []int{1, 1, 1},
			expected: []int64{3, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := distance(tt.nums)
			if len(result) != len(tt.expected) {
				t.Fatalf("length mismatch: got %d, want %d", len(result), len(tt.expected))
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("at index %d, got %d, want %d", i, v, tt.expected[i])
				}
			}
		})
	}
}
