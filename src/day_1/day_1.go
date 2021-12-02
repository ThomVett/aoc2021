// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ThomVett/aoc2021/src/utils"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReadData() string {
	dat, _ := os.ReadFile("data/day_1.txt")
	return string(dat)
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadIntsFromFile(file_path string) ([]int, error) {
	r, _ := os.Open(file_path)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ComputeNumberOfIncreases(numbers []int) int {
	increase_counter := 0

	for idx, number := range numbers {
		if idx+1 >= len(numbers) {
			return increase_counter
		}
		if number < numbers[idx+1] {
			increase_counter = increase_counter + 1
		}
	}
	return increase_counter
}

func ComputeNumberOfIncreasesFromFile() int {
	numbers, _ := ReadIntsFromFile("data/day_1.txt")

	return ComputeNumberOfIncreases(numbers)
}

func ComputeNumberOfIncreasesSlidingWindow(window_length int) int {
	numbers, _ := ReadIntsFromFile("data/day_1.txt")

	sliding_window := make([]int, 0)
	for idx, _ := range numbers {
		if idx+window_length > len(numbers) {
			continue
		}
		sliding_window = append(sliding_window, utils.Sum(numbers[idx:idx+window_length]))
		fmt.Println(numbers[idx:idx+window_length], utils.Sum(numbers[idx:idx+window_length]))
	}

	return ComputeNumberOfIncreases(sliding_window)
}
