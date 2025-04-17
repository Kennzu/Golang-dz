package main

import "fmt"

func main() {
	const USDEUR = 0.88
	const USDRUB = 85
	result := USDRUB / USDEUR

	fmt.Printf("1 EUR = %.2f RUB", result)
}
