package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	data_path := filepath.Join(cwd, "day2", "data.txt")

	file, err := os.Open(data_path)
	if err != nil {
		fmt.Println(err)
	}
	part_one(file)
}

func part_one(file *os.File) {
	scanner := bufio.NewScanner(file)

	var safe_count uint = 0
	var index uint = 0

	for scanner.Scan() {
		var line string = scanner.Text()
		var numbers []string = strings.Split(line, " ")
		var prev = 0
		var incr = 0
		var safe bool = true
		fmt.Println("---NEW LEVEL---")
		for i, str_num := range numbers {
			if i == 0 {
				prev, _ = strconv.Atoi(str_num)
				continue
			}

			curr, _ := strconv.Atoi(str_num)
			fmt.Println("PREV:", prev, "CURR", curr)
			// Checking if increasing or decreasing
			// Removing no adjacent change

			if prev == curr {
				safe = false
				break
			} else if prev < curr && incr == 0 {
				incr = 1
			} else if prev > curr && incr == 0 {
				incr = -1
			}
			fmt.Println(prev, curr)
			// Checking diff is between 1-3
			var diff = curr - prev

			if 1 <= intAbs(diff) && 3 >= intAbs(diff) {
				if (diff > 0 && incr > 0) || (diff < 0 && incr < 0) {
					prev = curr
					continue
				} else {
					safe = false
					break
				}
			} else {
				safe = false
				break
			}
		}
		if safe {
			fmt.Println("Safe")
			safe_count++
		} else {
			fmt.Println("Unsafe")
		}
		index++
	}
	fmt.Println("Safe Count", safe_count)
}

func intAbs(value int) int {
	return int(math.Abs(float64(value)))

}
