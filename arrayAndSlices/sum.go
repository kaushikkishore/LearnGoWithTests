package main

func Sum(numbers []int) int {
	// return 15
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// return []int{Sum(numbers), Sum(numbers2)}
	// return nil
	/*
		lengthOfNumbers := len(numbersToSum)
		sum := make([]int, lengthOfNumbers)

		for i, numbers := range numbersToSum {
			sum[i] = Sum(numbers)
		}

		return sum
	*/

	var sums []int

	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTrails(numbersToSum ...[]int) []int {
	var sums []int

	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			trail := number[1:]
			sums = append(sums, Sum(trail))
		}
	}
	return sums
}
