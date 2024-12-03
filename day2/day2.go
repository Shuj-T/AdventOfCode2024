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
	fmt.Println(file)
	// var arr = []int{1, 2, 3, 4}
	// fmt.Println(RemoveIndex(arr, 0))
	// fmt.Println(RemoveIndex(arr, 1))
	// fmt.Println(RemoveIndex(arr, 2))
	// fmt.Println(RemoveIndex(arr, 3))
	// fmt.Println(arr)

	// part_one(file)
	part_two(file)

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

func part_two(file *os.File) {
	scanner := bufio.NewScanner(file)

	var safe_count uint = 0
	var index uint = 0

	for scanner.Scan() {
		var line string = scanner.Text()
		var numbers []string = strings.Split(line, " ")
		var prev = 0
		var incr = 0
		var safe bool = true

		var report []int = make([]int, len(numbers))
		var bad_index int = -1
		for i, str_num := range numbers {
			if i == 0 {
				prev, _ = strconv.Atoi(str_num)
				report[0] = prev
				continue
			}

			curr, _ := strconv.Atoi(str_num)
			report[i] = curr
			if !safe {
				continue
			}

			// Checking if increasing or decreasing
			// Removing no adjacent change

			if prev == curr {
				safe = false
				bad_index = i
				continue
			} else if prev < curr && incr == 0 {
				incr = 1
			} else if prev > curr && incr == 0 {
				incr = -1
			}

			// Checking diff is between 1-3
			var diff = curr - prev

			if 1 <= intAbs(diff) && 3 >= intAbs(diff) {
				if (diff > 0 && incr > 0) || (diff < 0 && incr < 0) {
					prev = curr
					continue
				} else {
					safe = false
					bad_index = i
					continue
				}
			} else {
				safe = false
				bad_index = i
				continue
			}
		}
		if safe {
			fmt.Println("SAFE", report)
			safe_count++
		} else {
			fmt.Println("Bad Index", bad_index)
			safe_damp := false
			for i, _ := range report {
				fmt.Println(report, "REMOVING LEVEL", i+1)
				if is_safe(RemoveIndex(report, i)) {
					safe_damp = true
					break
				}
			}
			if safe_damp {
				fmt.Println("Problem Damper Safe", report)
				safe_count++
			} else {
				fmt.Println("Unsafe", report)
			}
		}

		index++
	}
	fmt.Println("Safe Count", safe_count)
}

func intAbs(value int) int {
	return int(math.Abs(float64(value)))

}

func is_safe(levels []int) bool {
	fmt.Println("is_safe INPUT", levels)
	var prev = 0
	var incr = 0
	for i, level := range levels {
		if i == 0 {
			prev = level
			continue
		}

		curr := level
		// Checking if increasing or decreasing
		// Removing no adjacent change

		if prev == curr {
			return false
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
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func RemoveIndex(arr []int, index int) []int {
	var new_arr []int
	new_arr = append(new_arr, arr[:index]...)
	return append(new_arr, arr[index+1:]...)
}
