package pe

func IsMultipleOf(divisor int, dividend int) bool {
	return dividend%divisor == 0
}

func SumDivisibleBy(below int, by int) int {
	target := (below - 1) / by
	return by * target * (target + 1) / 2
}

func SumMultiples(below int, multiples []int) int {
	combinedMultiple := 1
	for _, m := range multiples {
		combinedMultiple *= m
	}

	sum := 0
	for _, m := range multiples {
		sum += SumDivisibleBy(below, m)
	}

	return sum - SumDivisibleBy(below, combinedMultiple)
}

func SumFibonacci(max int) int {
	current := 0
	prev1 := 1
	prev2 := 1

	sum := 0

	for prev1 < max {
		if current%2 == 0 {
			sum += current
		}

		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return sum
}
