package tools

// 反转Bytearray
func ReverseBytes(res []byte) []byte {
	for from, to := 0, len(res)-1; from < to; from, to = from+1, to-1 {
		res[from], res[to] = res[to], res[from]
	}
	return res
}
