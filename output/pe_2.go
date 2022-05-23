package output

import (
	"fmt"
	"strconv"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

func PE2(args []string) {
	max, _ := strconv.Atoi(args[2])

	fmt.Print(pe.SumEvenFibonacci(max))
}
