package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput(msg string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(msg + " ")
	val, err := reader.ReadString('\n')
	val = strings.Replace(val, "\n", "", -1)

	return val, err
}

func GetIntegerInput(msg string) (int, error) {
	if msg == "" {
		msg = "Input number for operation:"
	}
	val, err := GetInput(msg)

	if err != nil {
		return -1, err
	}

	iVal, err := strconv.Atoi(val)

	return iVal, err
}

func OutputError(err error) {
	log.New(os.Stderr, "", 1)
	log.Println(err)
}
