package leetcode

func shellSort(nums []int) []int {

	for gap := len(nums) >> 1; gap >= 1; gap >>= 1 {

		for i, l := gap, len(nums); i < l; i++ {

			j := i - gap
			temp := nums[i]
			for ; j >= 0; j -= gap {

				if nums[j] > temp {
					nums[j+gap] = nums[j]
				} else {
					break
				}
			}

			nums[j+gap] = temp
		}
	}

	return nums
}