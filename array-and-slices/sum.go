package main

func Sum(numbers []int) int {
	sum := 0

	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for i := range numbers {
		sums = append(sums, Sum(numbers[i]))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) > 0 {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}

	return sums
}
