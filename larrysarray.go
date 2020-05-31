package HackerRankSolutions

//This functions will evaluate whether the array can be sorted using the following operation any number of times:
//Choose any  consecutive indices and rotate their elements in such a way that ABC->BCA->CAB->ABC
func larrysArray(A []int32) string {
	inversionCount := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				inversionCount++
			}
		}
	}
	if (len(A)%2 != 0 && inversionCount%2 == 0) || (len(A)%2 == 0 && inversionCount%2 == 0) {
		return "YES"
	}
	return "NO"
}
