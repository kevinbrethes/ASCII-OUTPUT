package errors

import (
	"fmt"
)

func noArguments(args []string) bool { //handling "no arguments" error
	if len(args) == 0 {
		return true
	}
	return false
}

func fewArguments(args []string) bool { //handling "few arguments" error
	if len(args) < 3 {
		return true
	}
	return false
}

func manyArguments(args []string) bool {
	if len(args) > 3 {
		return true
	}
	return false
}

func invalidCharacter(args []string) bool { //handling "non-valid character" error
	for _, val := range args {
		for i := range val {
			if val[i] < 32 || val[i] > 126 {
				return true
			}
		}
	}
	return false
}

func outputError(args []string) bool { //handling invalid argument
	if len(args[2]) < 10 {
		return true
	}
	if args[2][0:9] != "--output=" {
		return true
	}
	return false
}

func HandlingError(args []string) bool {
	switch {
	case noArguments(args):
		fmt.Println("You didn't give any arguments !")
		return true
	case invalidCharacter(args):
		fmt.Println("You used a non-valid character.")
		return true
	case manyArguments(args):
		fmt.Println("You have given too many arguments !")
		return true
	case fewArguments(args):
		fmt.Println("Too few arguments !")
		return true
	case outputError(args):
		fmt.Println("Invalid output arguments.")
		fmt.Println("Syntax : <text> <font> --output=<filename.txt>")
		return true
	default:
		return false
	}
}
