package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	l := len(nums)
	sort.Ints(nums)

	res := make([][]int, 0)

	for a := 0; a < l; {

		b := a + 1

		c := l - 1

		for b < c {

			if -nums[a] >= nums[b]+nums[c] {

				if -nums[a] == nums[b]+nums[c] {
					res = append(res, []int{nums[a], nums[b], nums[c]})
				}

				b++
				for b < l && nums[b] == nums[b-1] {
					b++
				}

			} else {

				c--
				for c >= 1 && nums[c] == nums[c+1] {
					c--
				}
			}
		}

		a++
		for a < l && nums[a] == nums[a-1] {
			a++
		}
	}

	return res
}

func main() {

	nums := []int{-2,0,1,1,2}
	fmt.Println(threeSum(nums))
}
