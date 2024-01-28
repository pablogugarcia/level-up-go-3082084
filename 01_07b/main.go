package main

import (
	"flag"
	"log"
	"strings"
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stack := make([]string, 0)

	e := strings.Split(expr, "")

	for _, v := range e {
		if v == "{" || v == "[" || v == "(" {

		}
		if v == "]" {
			if stack[len(stack)-1] != "]" {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if v == "}" {
			if stack[len(stack)-1] != "]" {
				return false
			}
			stack = stack[:len(stack)-1]
		}

		if v == ")" {
			if stack[len(stack)-1] != "]" {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
