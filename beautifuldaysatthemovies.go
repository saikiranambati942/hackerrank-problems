package HackerRankSolutions

import (
	"math"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	var count int32
	for m := i; m <= j; m++ {
		diff := math.Abs(float64(m - reverse(m)))
		if int32(diff)%k == 0 {
			count++
		}
	}
	return count
}

func reverse(n int32) int32 {
	var revNum int32 = 0
	for n > 0 {
		revNum = revNum*10 + n%10
		n = n / 10
	}
	return revNum
}
