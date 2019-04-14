package util

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
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

func GetFloatInput(msg string) (float64, error) {
	if msg == "" {
		msg = "Input number of operation:"
	}

	val, err := GetInput(msg)

	if err != nil {
		return float64(-1), err
	}

	fVal, err := strconv.ParseFloat(val, 64)
	return fVal, err
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

func StringArrayBinarySearch(a []string, x string) (int, error) {
	l, r := 0, len(a)

	for l <= r {
		m := int(math.Floor(float64(l + (r-l)/2)))

		comparison := strings.Compare(a[m], x)

		switch comparison {
		case 0:
			return m, nil
		case -1:
			l = m + 1
			break
		case 1:
			r = m - 1
		}
	}

	return -1, errors.New("element not in array")
}
