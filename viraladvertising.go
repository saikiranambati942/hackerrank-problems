package HackerRankSolutions

func viralAdvertising(n int32) int32 {
	var (
		liked      int32 = 2
		recipients int32
		cumulative int32 = 2
	)
	y := int(n)
	for i := 2; i <= y; i++ {
		recipients = liked * 3
		liked = recipients / 2
		cumulative = cumulative + liked
	}
	return cumulative

}
