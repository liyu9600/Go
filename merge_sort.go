package main

func mergeSort(nums []int) []int {

	length := len(nums)
	if length == 1 {
		return nums
	}
	middle := length / 2

	return merge(mergeSort(nums[0:middle]), mergeSort(nums[middle:length]))
}

func merge(leftNums []int, rightNums []int) []int {

	leftNumsLength := len(leftNums)
	rightNumsLength := len(rightNums)

	var temp []int
	l, r := 0, 0
	for ; l < leftNumsLength && r < rightNumsLength; {
		if leftNums[l] <= rightNums[r] {
			temp = append(temp, leftNums[l])
			l++
		} else {
			temp = append(temp, rightNums[r])
			r++
		}
	}

	if l == leftNumsLength {
		temp = append(temp, rightNums[r: rightNumsLength]...)
	} else {
		temp = append(temp, leftNums[l: leftNumsLength]...)
	}

	return temp
}