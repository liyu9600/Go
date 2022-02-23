package leetcode


func sort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {

	if start >= end {
		return
	}

	p := partition(nums, start, end)
	quickSort(nums, start, p-1)
	quickSort(nums, p+1, end)
}

func partition(nums []int, start, end int) int {

	nums[(start+end)/2], nums[end] = nums[end], nums[(start+end)/2]
	i, j, p := start, start, nums[end]

	for ; j < end; j++ {

		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = p, nums[i]

	return i
}