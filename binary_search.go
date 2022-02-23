package leetcode

//查找第一个值等于给定值的元素
func bsearch(nums []int, target int) int {

	for low, high := 0, len(nums) - 1; low <= high; {

		mid := low + (high-low)>>1

		if nums[mid] < target  {
			low = mid + 1
		} else if nums[mid] > target  {
			high = mid - 1
		} else {

			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			high = mid - 1
		}
	}

	return -1
}

//查找最后一个值等于给定值的元素
func bsearch2(nums []int, target int) int {

	for low, high := 0, len(nums)-1; low <= high; {

		mid := low + (high-low)>>1

		if nums[mid] < target  {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {

			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}

//查找第一个大于等于给定值的元素
func bsearch3(nums []int, target int) int {

	for low, high := 0, len(nums) - 1; low <= high; {

		mid := low + (high-low)>>1

		if nums[mid] >= target  {

			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

//查找最后一个小于等于给定值的元素
func bsearch4(nums []int, target int) int {

	for low, high := 0, len(nums) - 1; low <= high; {

		mid := low + (high-low)>>1

		if nums[mid] <= target {

			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}

			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
