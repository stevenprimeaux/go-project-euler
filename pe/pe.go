package pe

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

func EvenFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 2
	}
	return 4*EvenFibonacci(n-1) + EvenFibonacci(n-2)
}

func SumEvenFibonacci(max int) int {
	sum := 0
	next := 0
	for i := 0; true; i++ {
		next = EvenFibonacci(i)
		if next > max {
			break
		}
		sum += next
	}
	return sum
}
