package maxareaofisland

func maxAreaOfIsland(grid [][]int) int {

	maxArea := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				area := islandArea(grid, i, j)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func islandArea(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	sumOfArea := grid[i][j] // 現在位置の面積を足す
	grid[i][j] = 0          // 現在位置を到達済み（海扱い）とする
	// 上下左右の池面の面積を足す（grid上に領域が存在する場合のみ）
	if 0 < i {
		sumOfArea += islandArea(grid, i-1, j)
	}
	if i+1 < len(grid) {
		sumOfArea += islandArea(grid, i+1, j)
	}
	if 0 < j {
		sumOfArea += islandArea(grid, i, j-1)
	}
	if j+1 < len(grid[i]) {
		sumOfArea += islandArea(grid, i, j+1)
	}

	return sumOfArea
}
