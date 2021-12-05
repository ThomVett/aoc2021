package utils

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func SumBool(array []bool) int {
	result := 0
	for _, v := range array {
		if v {
			result += 1
		}
	}
	return result
}
