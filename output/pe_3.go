package output

import (
	"fmt"
	"strconv"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

func PE3(args []string) {
	n, _ := strconv.Atoi(args[2])

	fmt.Print(pe.LargestPrimeFactor(n))
}
