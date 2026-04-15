package solution

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name 			string
		nums 			[]int
		expectedLen 	int
		expectedNums	[]int
	}{
		{
			name: "test1",
			nums: []int{1, 2, 2, 2, 3, 4, 4, 4, 5, 5, 6, 7, 7},
			expectedLen: 11,
			expectedNums: []int{1, 2, 2, 3, 4, 4, 5, 5, 6, 7, 7},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputNums := make([]int, len(tt.nums))
			copy(inputNums, tt.nums)
			
			actualLen := removeDuplicates(inputNums)
			
			if actualLen != tt.expectedLen {
				fmt.Println(tt.name, "error")
				fmt.Println("期望数组内容: ", tt.expectedNums)
				fmt.Println("实际数组内容: ", tt.nums)
				t.Errorf("期望长度 %d, 实际得到 %d", tt.expectedLen, actualLen)
			} else {
				fmt.Println(tt.name, "ok")
			}
		})
	}
}