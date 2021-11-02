package main

func Sum(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}
