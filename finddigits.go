package HackerRankSolutions

func findDigits(n int32) int32 {
	var count int32
	var y, x int32
	y = n
	x = n % 10
	for y > 0 {
		y = (y - x) / 10
		if x != 0 && n%x == 0 {
			count++
		}
		x = y % 10
	}
	return count
}
