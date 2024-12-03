package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	data_path := filepath.Join(cwd, "day3", "data.txt")

	file, err := os.Open(data_path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	part_two(file)
}

func part_one(file *os.File) {
	scanner := bufio.NewScanner(file)
	var match = []rune{'m', 'u', 'l', '(', ' ', ',', ' ', ')'}
	var i int = 0
	fmt.Println(match, i)
	var buffer bytes.Buffer
	var first_num = -1
	var second_num = -1
	var total = 0
	var digit_count int
	for scanner.Scan() {
		var line string = scanner.Text()
		fmt.Println(line)
		for _, char := range line {
			fmt.Println(string(match[i]), "==", string(char))
			if match[i] == ' ' && is_int(char) && digit_count <= 3 {
				// fmt.Println(string(char))
				buffer.WriteString(string(char))
				digit_count++
			} else if match[i] == ' ' && !is_int(char) {
				i++
				digit_count = 0
				if first_num == -1 {

					first_num, _ = strconv.Atoi(buffer.String())
				} else {
					second_num, _ = strconv.Atoi(buffer.String())
				}
				buffer.Reset()
			}

			if match[i] == char {
				i++
			} else if digit_count != 0 {
				continue
			} else {
				fmt.Println("BROKEN")
				i = 0
				digit_count = 0
				buffer.Reset()
				first_num = -1
			}
			if i >= len(match) {
				fmt.Println(first_num, "*", second_num)
				total += first_num * second_num
				first_num = -1
				second_num = -1
				i = 0
			}

		}

	}
	fmt.Println("Total", total)
}

func part_two(file *os.File) {
	scanner := bufio.NewScanner(file)
	var enabled = true
	var match = []rune{'m', 'u', 'l', '(', ' ', ',', ' ', ')'}
	var do = []rune{'d', 'o', '(', ')'}
	var dont = []rune{'d', 'o', 'n', '\'', 't', '(', ')'}
	var i int = 0
	var j int = 0
	fmt.Println(match, i)
	var buffer bytes.Buffer
	var first_num = -1
	var second_num = -1
	var total = 0
	var digit_count int
	for scanner.Scan() {
		var line string = scanner.Text()
		fmt.Println(line)
		for _, char := range line {
			if (j < len(do) && do[j] == char) || (j < len(dont) && dont[j] == char) {
				fmt.Println(string(dont[i]), "===", string(char))
				j++
				if char == ')' && len(do) == j {
					enabled = true
					fmt.Println("ENABLED")
				} else if char == ')' && len(dont) == j {
					enabled = false
					fmt.Println("DISABLED")

				}
				continue
			} else {
				j = 0
			}
			fmt.Println(string(match[i]), "==", string(char))

			if match[i] == ' ' && is_int(char) && digit_count <= 3 {
				// fmt.Println(string(char))
				buffer.WriteString(string(char))
				digit_count++
			} else if match[i] == ' ' && !is_int(char) {
				i++
				digit_count = 0
				if first_num == -1 {

					first_num, _ = strconv.Atoi(buffer.String())
				} else {
					second_num, _ = strconv.Atoi(buffer.String())
				}
				buffer.Reset()
			}

			if match[i] == char {
				i++
			} else if digit_count != 0 {
				continue
			} else {
				fmt.Println("BROKEN")
				i = 0
				digit_count = 0
				buffer.Reset()
				first_num = -1
			}
			if i >= len(match) {
				fmt.Println(first_num, "*", second_num)
				if enabled {
					total += first_num * second_num
				}
				first_num = -1
				second_num = -1
				i = 0
			}

		}

	}
	fmt.Println("Total", total)
}

func is_int(str_value rune) bool {
	_, err := strconv.Atoi(string(str_value))
	if err != nil {
		return false
	}
	return true
}
