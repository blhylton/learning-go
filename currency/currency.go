package currency

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learningGo/util"
	"net/http"
	"os"
	"strings"
)

//This would be a const if constant arrays were possible
var currencies = [...]string{"AUD", "BGN", "BRL", "CAD", "CHF", "CNY", "CZK", "DKK", "GBP", "HKD", "HRK", "HUF", "IDR", "ILS", "INR",
	"ISK", "JPY", "KRW", "MXN", "MYR", "NOK", "NZD", "PHP", "PLN", "RON", "RUB", "SEK", "SGD", "THB", "TRY", "USD", "ZAR"}

func Run() {
	fmt.Println("===== Currency Conversion =====")
	fmt.Println("At any prompt, enter ? for help or q to quit")

	base := GetCurrencyValue("Enter base currency:")
	if base == "q" {
		return
	}

	res := GetCurrencyValue("Enter desired currency result:")
	if res == "q" {
		return
	}

	if val, err := util.GetFloatInput("Enter initial currency value:"); err == nil {
		rate := GetExchangeRate(base, res)
		fmt.Printf("%v %s is equivalent to %v %s\n", val, base, val*rate, res)
		return
	} else {
		util.OutputError(err)
		return
	}
}

func PrintCurrencyHelp() {
	currPage, pageSize, i := 1, 10, 0

	for i < len(currencies) {
		var currencySlice []string
		if currPage*pageSize > len(currencies) {
			currencySlice = currencies[i+(currPage-1*pageSize):]
		} else {
			currencySlice = currencies[i+(currPage-1*pageSize) : pageSize*currPage]
		}
		for _, elem := range currencySlice {
			fmt.Println(elem)
			fmt.Println("Press enter to continue...")
			_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
		currPage++
		i += currPage * pageSize
	}
}

func GetCurrencyValue(msg string) string {
	for {
		base, err := util.GetInput(msg)

		if err != nil {
			util.OutputError(err)
		}

		if base == "q" {
			return "q"
		}

		if base == "?" {
			PrintCurrencyHelp()
		}

		base = strings.ToUpper(base)

		_, locErr := util.StringArrayBinarySearch(currencies[0:], base)

		if locErr == nil {
			return base
		}

		util.OutputError(locErr)
	}
}

func GetExchangeRate(base string, res string) float64 {
	response, requestErr := http.Get("https://api.exchangeratesapi.io/latest/?base=" + base + "&symbols=" + res)

	if requestErr != nil {
		util.OutputError(requestErr)
		return 0
	}

	defer response.Body.Close()

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		util.OutputError(readErr)
		return 0
	}

	var exchangeData map[string]interface{}
	parseErr := json.Unmarshal(body, &exchangeData)

	if parseErr != nil {
		util.OutputError(parseErr)
		return 0
	}

	rates := exchangeData["rates"].(map[string]interface{})

	fRate := rates[res].(float64)
	return fRate
}
