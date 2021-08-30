package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func calculate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	number1, _ := strconv.Atoi(vars["a"])
	number2, _ := strconv.Atoi(vars["b"])
	var result string
	calculator := calculator{number1, number2}
	switch vars["op"] {
	case "sum":
		result = calculator.sum()
		save("Sum: ", result)
	case "sub":
		result = calculator.subtract()
		save("Subtract: ", result)
	case "mul":
		result = calculator.multiply()
		save("Multiply: ", result)
	case "div":
		result = calculator.division()
		save("Division: ", result)
	default:
		fmt.Fprint(w, "This is a invalid operation!")
	}
	fmt.Fprintf(w, result)
}

func history(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Historics: %v\n", getHistory())
}

func main() {
	fmt.Print("Serving at http://localhost:8080")
	r := mux.NewRouter()
	r.HandleFunc("/calc/{op}/{a}/{b}", calculate)
	r.HandleFunc("/calc/history", history)
	http.ListenAndServe(":8080", r)
}
