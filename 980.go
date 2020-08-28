package main

func uniquePathsIII(grid [][]int) int {
    startX := 0
    startY := 0
    num := 0
    for i := 0; i < len(grid); i ++ {
        for j := 0; j < len(grid[0]); j ++ {
            if (grid[i][j] == 1) {
                startX = i
                startY = j
            } else if grid[i][j] != -1 {
                num ++
            }
        }
    }
    return dfs(grid, startX, startY, num)
}

func dfs(grid [][]int, x int, y int, num int) int {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == -1 {
        return 0
    }
    if grid[x][y] == 2 {
        if num == 0 {
            return 1
		}
		
        return 0
        
    }
    grid[x][y] = -1
    ret := 0
    ret += dfs(grid, x, y - 1, num - 1)
    ret += dfs(grid, x, y + 1, num - 1)
    ret += dfs(grid, x - 1, y, num - 1)
    ret += dfs(grid, x + 1, y, num - 1)
    grid[x][y] = 0
    return ret
}