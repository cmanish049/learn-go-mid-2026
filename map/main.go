package main

import "fmt"

func main() {
	currencyCode := make(map[string]string)

	fmt.Println(currencyCode)

	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"
	fmt.Println("currencyCode map contents:", currencyCode)

	newCurrencyCode := map[string]string{
		"USD": "US Dollar",
		"GBP": "Pound Sterling",
		"EUR": "Euro",
	}
	newCurrencyCode["INR"] = "Indian Rupee"
	fmt.Println("currencyCode map contents:", newCurrencyCode)

	delete(newCurrencyCode, "INR")

	fmt.Println("map after deletion", newCurrencyCode)

	cyCode := "INR"
	if currencyName, ok := newCurrencyCode[cyCode]; ok {
		fmt.Println("Currency name for currency code", cyCode, " is", currencyName)
		return
	}
	fmt.Println("Currency name for currency code", cyCode, "not found")

}
