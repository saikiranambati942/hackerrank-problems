package HackerRankSolutions

func angryProfessor(k int32, a []int32) string {
	var count int32 = 0
	for _, v := range a {
		if v <= 0 {
			count++
		}
	}
	if count >= k {
		return "NO"
	}
	return "YES"
}
