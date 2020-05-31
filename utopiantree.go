package HackerRankSolutions

func utopianTree(n int32) int32 {
	var result int32
	for i := 0; i <= int(n); i++ {
		result = result*2 + 1
		i++
	}
	if n%2 == 0 {
		return result
	}
	return result * 2
}
