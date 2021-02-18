package main

func minKBitFlips(A []int, K int) int {
	var queue []int
	l := len(A)
	res := 0
	for i := 0; i < l; i++ {
		if len(queue) > 0 && i >= queue[0]+K {
			queue = queue[1:]
		}
		if len(queue)%2 == A[i] {
			if i+K > l {
				return -1
			}
			queue = append(queue, i)
			res++
		}
	}

	return res
}
