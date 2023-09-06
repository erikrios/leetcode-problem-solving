package main

import "fmt"

func main() {
	fmt.Println(
		calculateTax(
			[][]int{
				{3, 50},
				{7, 10},
				{12, 25},
			},
			10,
		),
	)

	fmt.Println(
		calculateTax(
			[][]int{
				{1, 0},
				{4, 25},
				{5, 50},
			},
			2,
		),
	)

	fmt.Println(
		calculateTax(
			[][]int{
				{2, 50},
			},
			0,
		),
	)
}

func calculateTax(brackets [][]int, income int) float64 {
	var tax float64

	for i := 0; i < len(brackets); i++ {
		if income == 0 {
			break
		}

		var taxBracket int

		bracket := brackets[i]

		if i == 0 {
			taxBracket = bracket[0]
		} else {
			prevBracket := brackets[i-1]
			taxBracket = bracket[0] - prevBracket[0]
		}

		if taxBracket >= income {
			taxBracket = income
			income = 0
		} else {
			income -= taxBracket
		}

		taxPercentage := bracket[1]

		tax += (float64(taxBracket) * float64(taxPercentage) / 100.0)
	}

	return tax
}
