package main

func SumAll(numbersToSum ...[]int) (sums []int) {
	// for _, arr := range numbersToSum {
	// 	var sum = 0
	// 	for _, n := range arr {
	// 		sum += n
	// 	}
	// 	sums = append(sums, sum)
	// }
	// return sums

	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}
	return sums
}
