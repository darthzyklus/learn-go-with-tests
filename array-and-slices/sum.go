package main

func Sum(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAllTails(groupOfNumbers ...[]int) []int {
	var sums []int

	for _, group := range groupOfNumbers {
		if len(group) == 0 {
			sums = append(sums, 0)
		} else {
			tail := group[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
