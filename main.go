package main

import (
	"fmt"
	"math"
	"newprac/pkg/formulars"
	"os"
)

func main() {
	input, err := formulars.ReadDataFile()
	if err != nil {
		fmt.Println("error: reading error,", err)
		os.Exit(0)
	}
	result1 := math.Round(formulars.Mean(input))

	fmt.Printf("Avarage is: %d\n", int(result1))
}
