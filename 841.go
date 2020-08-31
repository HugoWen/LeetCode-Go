package main

func canVisitAllRooms(rooms [][]int) bool {
    visited := make(map[int]bool)
    visited[0] = true
    flag := false
    var dfs func(int)
    dfs = func(pos int) {
        // fmt.Println(pos, visited, flag)
        if allVisited(visited, len(rooms)) == true {
            flag = true
        }
        
        for i := 0; i < len(rooms[pos]); i ++ {
            if visited[rooms[pos][i]] == true {
                if rooms[pos][i] == 0 {
                    if allVisited(visited, len(rooms)) == true {
                        flag = true
                        break
                    }
                }
                continue;
            }
            visited[rooms[pos][i]] = true
            dfs(rooms[pos][i])
            //visited[rooms[pos][i]] = false
        }
    }
    dfs(0)
    return flag
}

func allVisited(visited map[int]bool, lenght int) bool {
    if len(visited) != lenght {
        return false
    }
    hasNoVisited := false
    for _, v  := range visited {
        if v == false {
            hasNoVisited = true
            break;
        }
    }
    if hasNoVisited == false {
        return true
    }
    return false
}