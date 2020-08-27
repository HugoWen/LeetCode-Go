package main

import "sort"

func findItinerary(tickets [][]string) []string {
    var tourMap = map[string][]string{}
    var ret []string

    for _, ticket := range tickets {
        tourMap[ticket[0]] = append(tourMap[ticket[0]], ticket[1])
    }

    for key := range tourMap {
        sort.Strings(tourMap[key])
    }

    var dfs func(string)
    dfs = func(start string) {
        for {
            v, ok := tourMap[start]
            if ok == false || len(v) == 0 {
                break;
            }

            tmp := tourMap[start][0]
            tourMap[start] = tourMap[start][1:]
            dfs(tmp)
        }
        ret = append(ret, start)
    }

    dfs("JFK")
    for i := 0; i < len(ret) / 2; i ++ {
        t := ret[i]
        ret[i] = ret[len(ret) - i - 1]
        ret[len(ret) - i - 1] = t
    }
    return ret
}