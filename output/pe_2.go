package output

import (
	"fmt"
	"os"
	"strconv"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

func PE2(args []string) {
	max, _ := strconv.Atoi(os.Args[2])

	fmt.Print(pe.SumFibonacci(max))
}
