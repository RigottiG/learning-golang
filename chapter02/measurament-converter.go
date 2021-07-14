package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: measurament-converter <value> <unity>")
		os.Exit(1)
	}

	unityOrigin := os.Args[len(os.Args)-1]
	valuesOrigin := os.Args[1 : len(os.Args)-1]

	var unityTarget string

	if unityOrigin == "celsius" {
		unityTarget = "fahrenheit"
	} else if unityOrigin == "km" {
		unityTarget = "miles"
	} else {
		fmt.Printf("%s isn't a valid unity", unityOrigin)
		os.Exit(1)
	}

	for i, v := range valuesOrigin {
		valueOrigin, err := strconv.ParseFloat(v, 64)

		if err != nil {
			fmt.Printf("Value %s in the position %d isn't a valid number \n", v, i)
			os.Exit(1)
		}

		var valueTarget float64

		if unityTarget == "celsius" {
			valueTarget = valueOrigin*1.8 + 32
		} else {
			valueTarget = valueOrigin / 1.60934
		}

		fmt.Printf("%.1f %s = %.2f %s\n",
			valueOrigin, unityOrigin,
			valueTarget, unityTarget)
	}
}
