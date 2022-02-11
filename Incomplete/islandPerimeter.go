func islandPerimeter(grid [][]int) int {
    hor := len(grid)
    ver := len(grid[0])
    res := 0
    for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == 1 {
                res += 4
                if i > 0 && grid[i - 1][j] == 1 { res -= 1 }
                if j > 0 && grid[i][j - 1] == 1 { res -= 1 }
                if i + 1 < hor && grid[i + 1][j] == 1 { res -= 1 }
                if j + 1 < ver && grid[i][j + 1] == 1 { res -= 1 }
            }
        }
    }
    return res
}
