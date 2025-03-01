package leetcode

// O(n)

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := seen[complement]; found {
			return []int{index, i}
		}
		seen[num] = i
	}
	return nil
}

// O(n²)
// func TwoSum(nums []int, target int) []int {
// 	for i, num := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if num+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
