package main

func climbStairs(n int) int {

	/*if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	res := climbStairs(n-1) + climbStairs(n - 2)

	return res*/

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	res := 2
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		res = n1 + n2
		n1, n2 = n2, res
	}

	return res
}
