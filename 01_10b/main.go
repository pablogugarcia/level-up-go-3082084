package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return .0, err
	}
	return parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) float64 {
	ops := strings.Fields(expr)
	left, err := parseOperand(ops[0])
	if err != nil {
		log.Fatal("Error while parsing frist operand")
	}
	right, err := parseOperand(ops[2])
	if err != nil {
		log.Fatal("Error while parsing sec operand")
	}
	f, ok := operators[ops[1]]
	if !ok {
		log.Fatalf("Error parsing the operation type. Unknown: %s operator", ops[1])
	}
	result := f(left, right)
	return result
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result := calculate(*expr)
	log.Printf("%s = %.2f\n", *expr, result)
}
