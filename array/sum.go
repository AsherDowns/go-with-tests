package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
