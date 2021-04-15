package main

import (
	"fmt"
	"os"

	"github.com/kakivasantha/learning/problems"
	"github.com/kakivasantha/learning/test"
)

func main() {
	fmt.Println("welcome ", test.Name)
	arg := "nothing"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	switch arg {
	case "isDivisible":
		var value, divider int
		fmt.Println("value: ")
		fmt.Scan(&value)
		fmt.Println("divider: ")
		fmt.Scan(&divider)
		if problems.IsDivisible(value, divider) {
			fmt.Printf("%v can be divisible by %v\n", value, divider)
		} else {
			fmt.Printf("%v cannot be divisible by %v\n", value, divider)
		}
	case "nothing":
		fmt.Println("no commands to run")
	default:
		fmt.Println("command does not exist")
	}
	fmt.Println("see you ", test.Name)

}
