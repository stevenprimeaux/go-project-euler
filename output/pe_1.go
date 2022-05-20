package output

import (
	"fmt"
	"strconv"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

func PE1(args []string) {
	below, _ := strconv.Atoi(args[2])
	multiples := []int{}
	for _, x := range args[3:] {
		multiple, _ := strconv.Atoi(x)
		multiples = append(multiples, multiple)
	}

	fmt.Print(pe.SumMultiples(below, multiples))
}
