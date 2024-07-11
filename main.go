package main

import (
	"fmt"
	"os"

	"github.com/quocsi014/weather-cli-app/feature"
)

func main() {
	args := os.Args

	if len(args) < 2{
		fmt.Println("Try:")
		fmt.Println("\tlsc <country code>: list city of country")
		fmt.Println("\tcw <city>: view current weather")
	}

	command := args[1]

	//args

	switch command {
	case "lsc":
		if len(args) < 3 {
			fmt.Println("Missing country")
			break
		}
		feature.ListCity(args[2])
	case "cw":
		if len(args) < 3 {
			fmt.Println("Missing city")
			break
		}
	default:
		fmt.Println("Command does not exist")
		fmt.Println("Try:")
		fmt.Println("\tlsc <country code>: list city of country")
		fmt.Println("\tcw <city>: view current weather")
	}
}
