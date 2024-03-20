package main

// Here we are building a simple calculator
// according to the corresponding book
import (
	"fmt"
	"strconv"
)

func add(i, j int) int { return i + j }
func sub(i, j int) int { return i - j }
func mul(i, j int) int { return i * j }
func div(i, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"5", "+", "2"},
		{"5", "-", "2"},
		{"5", "*", "2"},
		{"5", "/", "2"},
		{"5", "%", "2"},
		{"five", "+", "two"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunk, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunk(p1, p2)
		fmt.Println(result)
	}
}
