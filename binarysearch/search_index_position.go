package binarysearch

func SearchInsert(nums []int, target int) int {
	lower, higher := 0, len(nums)-1
	for lower <= higher {
		middle := (lower + higher) / 2

		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			lower = middle + 1
		} else {
			higher = middle - 1
		}
	}
	return lower
}
