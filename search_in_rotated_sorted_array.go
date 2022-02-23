package leetcode

func search(nums []int, target int) int {

	l := len(nums)

	if l == 1 {

		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	for low, high := 0, l-1; low <= high; {

		mid := low + (high-low)>>1

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[low] {

			if nums[mid] > target && nums[low] <= target { // nums[low] <= target <= nums[mid]
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //后半部分有序

			if nums[mid] < target && nums[high] >= target { // nums[mid] < target <= nums[high]
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
