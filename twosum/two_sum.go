package leetcode

func TwoSum(nums []int, target int) []int {
	for index, num := range nums {
		if (num + nums[index+1]) == target {
			return []int{index, index + 1}
		}
	}

	return []int{0, 0}
}
