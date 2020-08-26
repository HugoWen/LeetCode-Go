package main

func letterCombinations(digits string) []string {
    phoneMap := map[string]string {
        "2" : "abc",
        "3" : "def",
        "4" : "ghi",
        "5" : "jkl",
        "6" : "mno",
        "7" : "pqrs",
        "8" : "tuv",
        "9" : "wxyz",
    }

    if (len(digits) == 0) {
        return []string{}
    }

    var ret []string
    var searchCombine func(int, string)
    searchCombine = func(start int, path string) {
        if start == len(digits) {
            ret = append(ret, path)
        } else {
            digit := string(digits[start])
            cur := phoneMap[digit]
            for i := 0; i < len(cur); i ++ {
                path = path + string(cur[i])
                searchCombine(start + 1, path)
                path = path[0:len(path) - 1]
            }
        }
    }

    searchCombine(0, "")
    return ret
}