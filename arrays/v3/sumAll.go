package main

func Sum(numbers []int) int {
	sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += number[i]
	//}

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	//sum := 0
	//
	//
	//for _, number := range numbers {
	//	sum += number
	//}
	lengthOfnumbers := len(numbersToSum)
	sums = make([]int, lengthOfnumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
