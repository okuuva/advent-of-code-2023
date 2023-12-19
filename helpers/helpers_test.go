package helpers

import (
	"fmt"
	"reflect"
	"testing"
)

var makeRangeTestCases = []struct {
	start, end int
	result     []int
}{
	{
		start:  0,
		end:    10,
		result: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		start:  10,
		end:    0,
		result: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		start:  0,
		end:    -10,
		result: []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
	},
	{
		start:  -10,
		end:    0,
		result: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0},
	},
	{
		start:  0,
		end:    0,
		result: []int{},
	},
}

func TestMakeRange(t *testing.T) {
	for _, tc := range makeRangeTestCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			got := MakeRange(tc.start, tc.end)
			if !reflect.DeepEqual(got, tc.result) {
				t.Fatalf("expected '%v' for start '%v', end '%v'; got '%v'", tc.result, tc.start, tc.end, got)
			}
		})
	}
}
