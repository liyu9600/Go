package main

func kthLargestNumber(nums []string, k int) string {
	if len(nums) == 1 {
		return nums[0]
	}
	return find(nums, k, 0, len(nums)-1)
}

func find(nums []string, k, start, end int) string {
	p := partition(nums, start, end)

	if p+1 == k {
		return nums[p]
	}

	if p+1 > k { //倒序往前找
		return find(nums, k, start, p-1)
	} else {
		return find(nums, k, p+1, end)
	}
}

func partition(nums []string, start, end int) int {

	nums[(start+end)/2], nums[end] = nums[end], nums[(start+end)/2]
	i, j, p := start, start, nums[end]

	for ; j < end; j++ {

		if compareGreaterThan(nums[j], p) {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[end] = p, nums[i]
	return i
}

func compareGreaterThan(left, right string) bool {
	if len(left) > len(right) {
		return true
	}

	if len(left) < len(right) {
		return false
	}

	for i, j := 0, len(left); i < j; i++ {
		if left[i] > right[i] {
			return true
		} else if left[i] < right[i] {
			return false
		}
	}

	return true
}
