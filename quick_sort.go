package leetcode

func sort(nums []int) []int {
	quick_sort(nums, 0, len(nums)-1)
	return nums
}

func quick_sort(nums []int, start, end int) {

	if start >= end {
		return
	}

	p := partition(nums, start, end)
	quick_sort(nums, start, p-1)
	quick_sort(nums, p+1, end)
}

func partition(nums []int, start, end int) int {

	i, j, p := start, start, nums[end]
	for ; j < end; j++  {

		if nums[j] < p{
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = p, nums[i]

	return i
}