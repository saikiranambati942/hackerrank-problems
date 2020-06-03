package HackerRankSolutions

import (
	"fmt"
	"math"
)

func serviceLane(width []int32, cases [][]int32) []int32 {
	result := make([]int32, 0)
	for i := 0; i < len(cases); i++ {
		j := cases[i][0]
		k := cases[i][1]
		fmt.Println("width is :", width[j:k+1])
		a := findMin(width[j : k+1])
		result = append(result, a)

	}
	return result
}
func findMin(a []int32) int32 {
	var tmp int32 = math.MaxInt32
	for _, v := range a {
		if v < tmp {
			tmp = v
		}
	}
	return tmp
}
