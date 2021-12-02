// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package day_1

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReadData() string {
	dat, _ := os.ReadFile("data/day_1.txt")
	return string(dat)
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
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

func ReadNumberFromFile() int {
	f, _ := os.Open("data/day_1.txt")
	numbers, _ := ReadInts(f)

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
