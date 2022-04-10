package main

func countingSort(nums []int) []int {

	//1.找出最大值,确定数据范围
	max := 0
	for i, j := 0, len(nums); i < j; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	//2.生成计数数组
	c := make([]int, max+1)
	for i, j := 0, len(nums); i < j; i++ {
		c[nums[i]]++
	}

	for i, j := 1, max+1; i < j; i++ {
		c[i] = c[i] + c[i-1]
	}

	//3.排序
	s := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		index := c[nums[i]]
		s[index-1] = nums[i]
		c[nums[i]]--
	}
	return s
}
