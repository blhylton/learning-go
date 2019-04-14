package main

import (
	"bufio"
	"errors"
	"fmt"
	"learningGo/currency"
	"learningGo/factorial"
	"learningGo/triangles"
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
	fmt.Println("\t[2] Currency Exchange")
	fmt.Println("\t[3] Print Triangle (Smallest to Largest)")
	fmt.Println("\t[4] Print Triangle (Largest to Smallest)")
	fmt.Println("\t[q] Quit")
}

func HandleSelection(selection string) {
	util.Clear()
	switch selection {
	case "1":
		factorial.Run()
		break
	case "2":
		currency.Run()
		break
	case "3":
		triangles.Run("ASC")
		break
	case "4":
		triangles.Run("DESC")
		break
	case "q":
		os.Exit(0)
	default:
		util.OutputError(errors.New("invalid selection"))
	}

	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
