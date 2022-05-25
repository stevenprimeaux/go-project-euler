package output

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stevenprimeaux/go-project-euler/pe"
)

type Int3 struct {
	Arg1 int `json:"arg1"`
	Arg2 int `json:"arg2"`
	Arg3 int `json:"arg3"`
}

func PE1(w http.ResponseWriter, r *http.Request) {
	var args Int3
	json.NewDecoder(r.Body).Decode(&args)

	fmt.Fprint(w, pe.SumMultiples(args.Arg1, []int{args.Arg2, args.Arg3}))
}

func PE2(w http.ResponseWriter, r *http.Request) {
	var args Int3
	json.NewDecoder(r.Body).Decode(&args)

	fmt.Fprint(w, pe.SumEvenFibonacci(args.Arg1))
}

func PE3(w http.ResponseWriter, r *http.Request) {
	var args Int3
	json.NewDecoder(r.Body).Decode(&args)

	fmt.Fprint(w, pe.LargestPrimeFactor(args.Arg1))
}
