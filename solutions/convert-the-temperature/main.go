package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

const (
	kelvin             = 273.15
	fahrenheitMultiply = 1.80
	fahrenheitSum      = 32.00
)

func convertTemperature(celsius float64) []float64 {
	return []float64{
		celsius + kelvin,
		celsius*fahrenheitMultiply + fahrenheitSum,
	}
}
