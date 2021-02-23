package main

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows <= 1 {
		return s
	}
	ret := ""
	sSlice := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		sSlice[i] = []string{}
	}
	flag := -1
	i := 0
	for _, v := range s {
		sSlice[i] = append(sSlice[i], string(v))
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	// fmt.Println(sSlice)
	for j := 0; j < len(sSlice); j++ {
		for k := 0; k < len(sSlice[j]); k++ {
			ret += string(sSlice[j][k])
		}
	}
	return ret
}
