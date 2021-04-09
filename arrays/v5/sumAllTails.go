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
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers[1:]))
	}
	return sums
}
