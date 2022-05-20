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
