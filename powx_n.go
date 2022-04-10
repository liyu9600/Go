package main

/**
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
*/
func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	var res, base float64 = 1, x

	if n < 0 {
		n = -n
		base = 1 / x
	}

	for n > 0 {

		if n&1 == 1 {
			res = res * base
		}

		base = base * base
		n >>= 1
	}

	return res
}

func myPow2(x float64, n int) float64 {

	if n < 0 {
		res := process(x, -n)
		return 1 / res
	} else {
		return process(x, n)
	}
}

func process(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	y := process(x*x, n>>1)

	if n&1 == 1 {
		return y * x
	}

	return y
}
