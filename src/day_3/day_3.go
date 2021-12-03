package day_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NoOP() int {
	return 1
}

func ReadData(filepath string) [][]string {
	r, _ := os.Open(filepath)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	output := [][]string{}

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")
		output = append(output, text)
	}

	return output
}

func ParseStringArrayToDecimal(str_array []string) int64 {
	binary := strings.Join(str_array[:], "")
	number, _ := strconv.ParseInt(binary, 2, 64)
	return number
}

func CreateEnptyStringArray(length int) []string {
	string_array := []string{}

	for i := 0; i < length; i++ {
		string_array = append(string_array, "0")
	}
	return string_array
}

func GetMostCommonLeastCommon(datum [][]string, if_equal int) ([]string, []string) {
	// Gets the most common item (0 or 1 at each item in the string)
	most_common := []int{}

	// Populate the most_common array to ensure that it can handle different length of the sequences
	for i := 0; i < len(datum[0]); i++ {
		most_common = append(most_common, 0)
	}

	for _, data := range datum {
		for i := 0; i < len(data); i++ {
			if data[i] == "0" {
				most_common[i] = most_common[i] + 1
			}
		}
	}

	mid_point := len(datum) / 2

	most_common_str := CreateEnptyStringArray(len(most_common))
	least_common_str := CreateEnptyStringArray(len(most_common))

	for idx, common := range most_common {
		if common > mid_point {
			most_common_str[idx] = "0"
			least_common_str[idx] = "1"
		} else if common < mid_point {
			most_common_str[idx] = "1"
			least_common_str[idx] = "0"
		} else {
			most_common_str[idx] = strconv.Itoa(if_equal)
			least_common_str[idx] = strconv.Itoa(1 - if_equal)
		}
	}
	return most_common_str, least_common_str
}

func ComputePowerConsumption(filepath string) int64 {
	datum := ReadData(filepath)

	most_common_str, least_common_str := GetMostCommonLeastCommon(datum, 1)

	fmt.Println(ParseStringArrayToDecimal(most_common_str[:]))
	fmt.Println(ParseStringArrayToDecimal(least_common_str[:]))

	return ParseStringArrayToDecimal(most_common_str[:]) * ParseStringArrayToDecimal(least_common_str[:])
}

func FilterArrays(datum [][]string, filter string, array_index int) [][]string {
	// recursively filter the arrays until we keep only one
	var fill_value int
	if filter == "most_common" {
		fill_value = 1
	} else {
		fill_value = 1
	}
	most_common_str, least_common_str := GetMostCommonLeastCommon(datum, fill_value)

	var filter_value string

	if filter == "most_common" {
		filter_value = most_common_str[array_index]
	} else {
		filter_value = least_common_str[array_index]
	}

	filtered_data := [][]string{}

	for _, data := range datum {
		if data[array_index] == filter_value {
			filtered_data = append(filtered_data, data)
		}
	}

	if len(filtered_data) == 1 {
		return filtered_data
	} else {
		return FilterArrays(filtered_data, filter, array_index+1)
	}
}

func ComputeOxygenRating(filepath string) int64 {
	datum := ReadData(filepath)

	oxygen_str := FilterArrays(datum, "most_common", 0)[0]
	co2_str := FilterArrays(datum, "least_common", 0)[0]

	oxygen := ParseStringArrayToDecimal(oxygen_str[:])
	co2 := ParseStringArrayToDecimal(co2_str[:])

	fmt.Println(oxygen, co2)
	return oxygen * co2
}
