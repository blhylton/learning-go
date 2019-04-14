package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
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

func Clear() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	val, ok := clear[runtime.GOOS]

	if ok {
		val()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
