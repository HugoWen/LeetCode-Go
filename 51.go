package main

import "strings"

func solveNQueens(n int) [][]string {
    col := map[int]bool{}
    pie := map[int]bool{}
    na := map[int]bool{}
    queens := make(map[int]int, n)
    for i := 0; i < n; i++ {
        queens[i] = -1
    }
    ret := []map[int]int{}

    var dfs func(int)
    dfs = func(i int) {
        if i >= n {
            tmp := make(map[int]int, n)
            for k, v := range queens {
                tmp[k] = v
            }
            ret = append(ret, tmp)
            return
        }

        for j := 0; j < n; j ++ {
            if col[j] {
                continue
            }
            if pie[i + j] {
                continue
            }
            if na[i - j] {
                continue
            }
            
            col[j] = true;
            pie[i + j] = true;
            na[i - j] = true;
            queens[i] = j 
            dfs(i + 1)
            queens[i] = -1
            delete(col, j)
            delete(pie, i + j)
            delete(na, i - j)
        }
    }

    dfs(0)
    res := make([][]string, len(ret))
    for k := 0; k < len(ret); k ++ {
        res[k] = []string{}
        for r := 0; r < len(ret[k]); r ++ {
            s := ""
            if ret[k][r] == 0 {
                s = "Q" + strings.Repeat(".", n - 1)
            } else {
                s += strings.Repeat(".", ret[k][r]) + "Q" + strings.Repeat(".", n - ret[k][r] - 1)
            }
            res[k] = append(res[k], s)
        }
    }
    return res
}