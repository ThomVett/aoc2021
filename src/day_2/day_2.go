package day_2

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Array struct {
	direction string
	value     int
}

func getValueField(v *Array, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func getDirectionField(v *Array, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func ReadDirections(file_path string) []Array {
	list_of_directions := make([]Array, 0)
	r, _ := os.Open(file_path)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		direction_value, _ := strconv.Atoi(text[1])
		direction := Array{text[0], direction_value}
		list_of_directions = append(list_of_directions, direction)
	}
	return list_of_directions
}

func ComputeSolution(file_path string) int {
	list_of_directions := ReadDirections(file_path)

	forward := 0
	depth := 0

	for _, direction_struct := range list_of_directions {
		direction := getDirectionField(&direction_struct, "direction")
		value := getValueField(&direction_struct, "value")

		if direction == "forward" {
			forward = forward + value
		}

		if direction == "up" {
			depth = depth - value
		}

		if direction == "down" {
			depth = depth + value
		}
	}

	return forward * depth
}

func ComputeSolutionWithAim(file_path string) int {
	list_of_directions := ReadDirections(file_path)

	forward := 0
	depth := 0
	aim := 0

	for _, direction_struct := range list_of_directions {
		direction := getDirectionField(&direction_struct, "direction")
		value := getValueField(&direction_struct, "value")

		if direction == "forward" {
			forward = forward + value
			depth = depth + aim*value
		}

		if direction == "up" {
			aim = aim - value
		}

		if direction == "down" {
			aim = aim + value
		}
	}
	return forward * depth
}
