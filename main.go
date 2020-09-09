package main

import (
	"fmt"

	"github.com/michaelsidharta/cubic-weight/constant"
	"github.com/michaelsidharta/cubic-weight/external"
	"github.com/michaelsidharta/cubic-weight/service"
)

func main() {
	api := external.Init(constant.ApiURL)
	calc := service.InitCalculator(api)
	result, err := calc.GetAverage(constant.CategoryFilter)
	if err != nil {
		fmt.Println("Encounter error", err)
	}
	fmt.Printf("Average cubic weight of %s is %.2f\n", constant.CategoryFilter, result)
}
