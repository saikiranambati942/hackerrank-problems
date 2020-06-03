package HackerRankSolutions

func taumBday(B int32, W int32, BC int32, WC int32, Z int32) int64 {
	var b, w, bc, wc, z int64
	b = int64(B)
	w = int64(W)
	bc = int64(BC)
	wc = int64(WC)
	z = int64(Z)
	if bc > z+wc && bc > wc {
		return (wc * w) + ((z + wc) * b)
	} else if wc > z+bc && wc > bc {
		return (w * (z + bc)) + (b * bc)
	}
	return (b * bc) + (w * wc)
}
