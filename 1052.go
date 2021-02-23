package main

func maxSatisfied(customers []int, grumpy []int, X int) int {
	l := len(customers)
	total := 0
	maxValue := 0
	curValue := 0
	for i := 0; i < l; i++ {
		total += customers[i] * (1 - grumpy[i])
		if i < X {
			curValue += customers[i] * grumpy[i]
			maxValue = curValue
		} else {
			curValue = curValue + customers[i]*grumpy[i] - customers[i-X]*grumpy[i-X]
			maxValue = max(maxValue, curValue)
		}
	}

	return maxValue + total
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
