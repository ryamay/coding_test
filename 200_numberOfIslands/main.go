package numberofislands

func numIsland(grid [][]byte) int {
	numOfIslands := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				numOfIslands++
				// 地続きの地面はすべて到達済みとする
				paintIsland(grid, i, j)
			}
		}
	}

	return numOfIslands
}

func paintIsland(grid [][]byte, i, j int) {
	if grid[i][j] != '1' {
		return // 移動先が未到達の地面でない場合は終了する
	}

	grid[i][j] = '*' // 到達済みとマークする

	// 上下左右もマークする
	if 0 < i {
		paintIsland(grid, i-1, j)
	}
	if i+1 < len(grid) {
		paintIsland(grid, i+1, j)
	}
	if 0 < j {
		paintIsland(grid, i, j-1)
	}
	if j+1 < len(grid[i]) {
		paintIsland(grid, i, j+1)
	}
}
