package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	marked := [][]bool{}
	for _, row := range grid {
		marked = append(marked, make([]bool, len(row)))
	}
	count := 0
	for i, r := range marked {
		for j := range r {
			if !marked[i][j] {
				if grid[i][j] == byte('0') {
					continue
				}
				fmt.Println("=================== begin dfs ========================= ")
				dfs(grid, i, j, marked)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int, marked [][]bool) {
	marked[i][j] = true
	if string(grid[i][j]) == "0" {
		return
	}
	// find grid[i][j]'s adjacent items
	rowLength := len(grid)
	columnLength := len(grid[0])
	if j > 0 {
		// left
		if column := j - 1; !marked[i][column] && grid[i][column] == byte('1') {
			dfs(grid, i, column, marked)
		}
	}
	if j < columnLength-1 {
		// right
		column := j + 1
		// fmt.Println("i= ", i, "\tcolumn= ", column, columnLength)
		if !marked[i][column] && grid[i][column] == byte('1') {
			dfs(grid, i, column, marked)
		}
	}
	if i > 0 {
		// top
		if row := i - 1; !marked[row][j] && grid[row][j] == byte('1') {
			dfs(grid, row, j, marked)
		}
	}
	if i < rowLength-1 {
		// bottom
		if row := i + 1; !marked[row][j] && grid[row][j] == byte('1') {
			dfs(grid, row, j, marked)
		}
	}
}

func main() {
	grid := [][]byte{}
	grid = append(grid, []byte{0x01, 0x01, 0x01, 0x01, 0x00})
	grid = append(grid, []byte{0x01, 0x01, 0x00, 0x01, 0x00})
	grid = append(grid, []byte{0x01, 0x01, 0x00, 0x00, 0x00})
	grid = append(grid, []byte{0x00, 0x00, 0x00, 0x00, 0x00})
	fmt.Println("count: ", numIslands(grid))
}
