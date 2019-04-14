package main

import (
	"bufio"
	"errors"
	"fmt"
	"learningGo/factorial"
	"learningGo/util"
	"os"
)

func main() {
	for true {
		OutputMenu()
		input, err := util.GetInput("Please make a selection:")

		if err != nil {
			util.OutputError(err)
			os.Exit(1)
		}

		HandleSelection(input)
	}
}

func OutputMenu() {
	util.Clear()
	fmt.Println("=== Menu ===")
	fmt.Println("\t[1] Factorial Finder")
	fmt.Println("\t[q] Quit")
}

func HandleSelection(selection string) {
	util.Clear()
	switch selection {
	case "1":
		factorial.Run()
		break
	case "q":
		os.Exit(0)
	default:
		util.OutputError(errors.New("invalid selection"))
	}

	fmt.Println("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
