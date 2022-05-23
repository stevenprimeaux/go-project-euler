package main

import (
	"os"

	"github.com/stevenprimeaux/go-project-euler/output"
)

var problems = map[string]interface{}{
	"1": output.PE1,
	"2": output.PE2,
}

func main() {
	problems[os.Args[1]].(func([]string))(os.Args)
}
