package HackerRankSolutions

import "strings"

func repeatedString(s string, n int64) int64 {
	countinaString := strings.Count(s, "a")
	length := len(s)
	repeatedStrings := int(n) / length
	countString := countinaString * repeatedStrings
	remainder := int(n) % length
	if remainder != 0 {
		extraString := int(n) - (repeatedStrings * length)
		ex := s[:extraString]
		countString += strings.Count(ex, "a")
	}
	return int64(countString)
}
