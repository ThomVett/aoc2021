package day_4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ThomVett/aoc2021/src/matrix"
)

func ParseIntsFromString(input_string string, sep string) []int {
	int_list := []int{}
	space := regexp.MustCompile(`\s+`)
	input_string = strings.TrimSpace(input_string)

	// There are some duplicate spaces in the middle
	input_string = space.ReplaceAllString(input_string, " ")
	chars := strings.Split(input_string, sep)

	for _, bingo_char := range chars {
		x, _ := strconv.Atoi(bingo_char)
		int_list = append(int_list, x)
	}
	return int_list
}

func ParseData(filepath string) ([]int, []matrix.IntMatrix2D) {
	bingo_list := []int{}
	bingo_cards := []matrix.IntMatrix2D{}

	r, _ := os.Open(filepath)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	line_counter := 0

	matrix_data := [][]int{}

	for scanner.Scan() {
		if line_counter == 0 {
			bingo_list = ParseIntsFromString(scanner.Text(), ",")
		} else if scanner.Text() == "" {
			if len(matrix_data[0]) != 0 {
				matrix_object := matrix.IntMatrix2D{
					Data: matrix_data,
					Name: "test",
				}
				bingo_cards = append(bingo_cards, matrix_object)
			}
		} else {
			int_list := ParseIntsFromString(scanner.Text(), " ")
			matrix_data[(line_counter-2)%5] = int_list
		}

		line_counter = line_counter + 1
	}
	return bingo_list, bingo_cards
}

func Part1(filepath string) int {
	bingo_list, bingo_cards := ParseData(filepath)

	for _, bingo_number := range bingo_list {
		for i, _ := range bingo_cards {
			bingo_cards[i].Contains(bingo_number)
			fmt.Println(bingo_cards[i].Selected)
		}

		for _, bingo_card := range bingo_cards {
			if bingo_card.WinningCol() || bingo_card.WinningRow() {
				return bingo_number
			}
		}
	}
	return 0
}
