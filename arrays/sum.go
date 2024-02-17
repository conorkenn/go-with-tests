package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	/*
		for i := 0; i < 5; i++ {
			sum += numbers[i]
		}
	*/
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tails := nums[1:]
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
