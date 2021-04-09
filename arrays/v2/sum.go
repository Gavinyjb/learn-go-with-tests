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
