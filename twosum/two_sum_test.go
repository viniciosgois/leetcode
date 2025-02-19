package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("should return the indices that the sum of them is the target", func(t *testing.T) {
		var nums = []int{20, 50, 50}
		var expected = []int{1, 2}

		output := TwoSum(nums, 100)

		if !reflect.DeepEqual(output, expected) {
			t.Errorf("output %d expected %d, given %v", output, expected, nums)
		}
	})
}
