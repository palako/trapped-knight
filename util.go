package main

func abs(x int) int {
	y := x >> 63
	return (x ^ y) - y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
