package HackerRankSolutions

import (
	"strings"
)

func designerPdfViewer(h []int32, word string) int32 {
	word = strings.ToUpper(word)
	var result, tmp int32
	for _, v := range word {
		value := v - 64
		if h[value-1] > tmp {
			tmp = h[value-1]
		}
	}
	result = tmp * int32(len(word))
	return result
}
