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

	for scanner.Scan() {

		var line string = scanner.Text()
		var values []string = strings.Split(line, "   ")

		left[index], err = strconv.Atoi(values[0])
		right[index], err = strconv.Atoi(values[1])
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
}
