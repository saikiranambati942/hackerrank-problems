package HackerRankSolutions

func hurdleRace(k int32, height []int32) int32 {
	var result int32
	for _, v := range height {
		if v > result {
			result = v
		}
	}
	if k >= result {
		return 0
	}
	result = result - k
	return result
}
