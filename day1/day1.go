package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	data_path := filepath.Join(cwd, "day1", "data.txt")

	file, err := os.Open(data_path)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var left [1000]int
	var right [1000]int

	var index uint = 0

	// Part Two
	right_occur := make(map[int]int)

	for scanner.Scan() {

		var line string = scanner.Text()
		var values []string = strings.Split(line, "   ")

		left[index], err = strconv.Atoi(values[0])
		right[index], err = strconv.Atoi(values[1])

		// Part Two
		if !does_exist(right_occur, right[index]) {
			right_occur[right[index]] = 0
		}
		right_occur[right[index]] += 1

		index++
	}
	lefts := left[:]
	rights := right[:]

	sort.Ints(lefts)
	sort.Ints(rights)
	total := 0
	for i := 0; i < 1000; i++ {
		l := lefts[i]
		r := rights[i]
		if l > r {
			total += l - r
		} else {
			total += r - l
		}
	}
	fmt.Println("Total distance:", total)

	// Part Two
	sim_score := 0
	for _, number := range lefts {
		sim_score += number * get_value_or_zero(right_occur, number)
	}
	fmt.Println("Similarity Score:", sim_score)

}

func does_exist(map_to_check map[int]int, value int) bool {
	_, ok := map_to_check[value]
	return ok
}

func get_value_or_zero(map_to_check map[int]int, value int) int {
	value, ok := map_to_check[value]
	if !ok {
		return 0
	}
	return value
}
