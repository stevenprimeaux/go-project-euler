package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

func main() {
	below, _ := strconv.Atoi(os.Args[1])
	multiples := []int{}
	for _, x := range os.Args[2:] {
		multiple, _ := strconv.Atoi(x)
		multiples = append(multiples, multiple)
	}

	fmt.Print(pe.SumMultiples(below, multiples))
}
