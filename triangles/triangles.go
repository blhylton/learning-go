package triangles

import (
	"fmt"
	"learningGo/util"
	"strings"
)

func Run(direction string) {
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}

	amount, err := util.GetIntegerInput("Input the desired length of the longest line:")

	if err != nil {
		util.OutputError(err)
		return
	}

	if direction == "ASC" {
		for i := 1; i <= amount; i++ {
			PrintStars(i)
		}
	} else {
		for i := amount; i > 0; i-- {
			PrintStars(i)
		}
	}
	return
}

func PrintStars(amount int) {
	fmt.Println(strings.Repeat("*", amount))
	return
}
