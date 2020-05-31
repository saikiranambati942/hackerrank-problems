package HackerRankSolutions

import "math"

//This functions finds the maximum number of integers we can select from an array such that the absolute difference between any two numbers is less than or equal to 1.
func pickingNumbers(a []int32) int32 {
	m := make(map[int32]int32)
	for _, v := range a {
		m[v] += 1
	}
	var result float64
	for k, _ := range m {
		x := m[k]
		y := m[k+1]
		result = math.Max(result, float64(x+y))
	}
	return int32(result)
}
