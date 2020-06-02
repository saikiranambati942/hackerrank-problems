package HackerRankSolutions

func beautifulTriplets(d int32, arr []int32) int32 {
	m := make(map[int32]int)
	for i, v := range arr {
		m[v] = i
	}
	var count int32
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		_, ok1 := m[x+d]
		_, ok2 := m[x+2*d]
		if ok1 && ok2 {
			count++
		}
	}
	return count
}
