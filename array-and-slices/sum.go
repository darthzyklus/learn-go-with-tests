package main

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(groupOfNumbers ...[]int) []int {
	var sum []int

	for _, group := range groupOfNumbers {
		sum = append(sum, Sum(group))

	}

	return sum
}
