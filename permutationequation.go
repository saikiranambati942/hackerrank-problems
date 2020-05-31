package HackerRankSolutions

import "fmt"

func main() {
	p := []int32{5, 2, 1, 3, 4}
	fmt.Println(permutationEquation(p))
}
func permutationEquation(p []int32) []int32 {
	result := make([]int32, 0)
	m := make(map[int32]int32)
	for i, v := range p {
		m[v] = int32(i + 1)
	}
	for x := 1; x <= len(p); x++ {
		y1 := m[int32(x)]
		y := m[y1]
		result = append(result, y)
	}

	return result
}

//func permutationEquation(p []int32) []int32 {
//	result := make([]int32, 0)
//	y1 := 0
//	for x := 1; x <= len(p); x++ {
//		for i, v := range p {
//			if x == int(v) {
//				y1 = i + 1
//				for i, v := range p {
//					if y1 == int(v) {
//						result = append(result, int32(i+1))
//					}
//				}
//			}
//
//		}
//
//	}
//	return result
//}
