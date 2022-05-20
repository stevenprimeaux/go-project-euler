package pe

func IsMultipleOf(divisor int, dividend int) bool {
	return dividend%divisor == 0
}

func SumMultiples(below int, multiples ...int) (sum int) {
	sum = 0
	for d := 1; d < below; d++ {
		for _, m := range multiples {
			if IsMultipleOf(m, d) {
				sum += d
				break
			}
		}
	}
	return
}
