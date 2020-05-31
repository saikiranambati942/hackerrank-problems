package HackerRankSolutions

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	length := len(a)
	result := make([]int32, len(queries))
	r := int(k) % length
	b := a[:length-r]
	a = a[length-r:]
	a = append(a, b ...)
	for i, v := range queries {
		result[i] = a[v]
	}
	return result
}
