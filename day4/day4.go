package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type position struct {
	x int
	y int
}

const GRID_SIZE = 10

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	data_path := filepath.Join(cwd, "day4", "data.txt")

	file, err := os.Open(data_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	part_one(file)
}

func part_one(file *os.File) {
	scanner := bufio.NewScanner(file)
	var grid [GRID_SIZE][GRID_SIZE]rune
	var y = 0
	var x_pos []position
	var xmax_count = 0
	for scanner.Scan() {
		var line string = scanner.Text()
		for x, char := range line {
			grid[y][x] = char
			if char == 'X' {
				x_pos = append(x_pos, position{x, y})
				fmt.Print(string(char))
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println("")
		y++
	}
	for _, pos := range x_pos {
		if check_xmas(pos, grid, 0) {
			xmax_count++
		}
	}
	fmt.Println(len(x_pos))
	fmt.Println(xmax_count)

}

func check_xmas(start position, grid [GRID_SIZE][GRID_SIZE]rune, index int) bool {
	if start.x == -1 && start.y == -1 {
		return false
	}
	var i = index
	var xmas = []rune{'M', 'A', 'S'}
	directions := []string{"U", "D", "L", "R", "UL", "UR", "DL", "DR"}

	for _, dir := range directions {
		var pos = check(start, dir, xmas[i], grid)
		if i+1 == len(xmas) || check_xmas(pos, grid, i+1) {
			return true
		}
	}
	return false
}
func check(curr position, direction string, letter rune, grid [GRID_SIZE][GRID_SIZE]rune) position {
	var pos position

	switch direction {
	case "U":
		pos = position{curr.x, curr.y + 1}
	case "D":
		pos = position{curr.x, curr.y - 1}
	case "L":
		pos = position{curr.x + 1, curr.y}
	case "R":
		pos = position{curr.x - 1, curr.y}
	case "UL":
		pos = position{curr.x + 1, curr.y + 1}
	case "UR":
		pos = position{curr.x - 1, curr.y + 1}
	case "DL":
		pos = position{curr.x - 1, curr.y - 1}
	case "DR":
		pos = position{curr.x + 1, curr.y - 1}
	}

	if get_from_grid(pos, grid) == letter {
		return pos
	}

	return position{-1, -1}
}

func get_from_grid(pos position, grid [GRID_SIZE][GRID_SIZE]rune) rune {
	if 0 <= pos.x && pos.x < len(grid[0]) && 0 <= pos.y && pos.y < len(grid) {
		return grid[pos.y][pos.x]
	}
	return ' '
}
