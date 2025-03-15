package binarysearch

func MaximumCountOfPositiveIntegerAndNegativeInteger(nums []int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	neg := left

	left, right = 0, len(nums)

	for left < right {
		mid := (left + right) / 2

		if nums[mid] <= 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pos := len(nums) - left

	if neg > pos {
		return neg
	}
	return pos
}
