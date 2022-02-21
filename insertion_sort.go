package leetcode

func insertionSort(nums []int) []int {

	for i, l := 1, len(nums); i < l; i++ {

		j := i - 1
		value := nums[i]
		for ; j >= 0; j-- {

			if value < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}

	return nums
}
