package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sum := 0
		if len(numbers) > 0 {
			tail := numbers[1:]
			sum = Sum(tail)
		}
		sums = append(sums, sum)
	}
	return
}
