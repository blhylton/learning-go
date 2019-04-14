package factorial

import (
	"errors"
	"fmt"
	"github.com/JohnCGriffin/overflow"
	"learningGo/util"
	"math/big"
	"strconv"
)

func Run() {
	val, err := util.GetIntegerInput("Input number to find factorial:")

	if err != nil {
		util.OutputError(err)
		return
	}

	if val < 0 {
		util.OutputError(errors.New("number can't be negative - factorial of a negative number is undefined"))
		return
	}

	if val == 0 {
		fmt.Printf("Factorial of %v is always %v\n", val, 0)
		return
	}

	factorial, ok := FindFactorial(val)

	if !ok {
		factorial = FindBigIntFactorial(val)
	}

	fmt.Printf("Factorial of %v is %v\n", val, factorial)
	return

}

func FindFactorial(input int) (string, bool) {
	sum := 1

	for i := 1; i <= int(input); i++ {
		val, ok := overflow.Mul(sum, i)

		if !ok {
			return "", ok
		}

		sum = val
	}

	return strconv.Itoa(sum), true
}

func FindBigIntFactorial(input int) string {
	sum := big.NewInt(1)
	sum = sum.MulRange(1, int64(input))
	return sum.Text(10)
}
