package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/quocsi014/weather-cli-app/feature"
)

func main() {
	args := os.Args
	envPath := os.Getenv("WEA_ENV")
	if envPath == "" {
		fmt.Println("\"EXPORT WEA_ENV=PATH/.wea\" to storage enviroment variable")
		return
	}
	err := godotenv.Load(envPath)
	if err != nil {
			log.Fatalf("Failed to load .env file: %v", err)
	}
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
		feature.ViewCurrentWeather(args[2])
	case "ckey":
		if len(args) < 3{
			fmt.Println("Missing api key")
			break
		}
		feature.ChangeKey(args[2])
	default:
		fmt.Println("Command does not exist")
		fmt.Println("Try:")
		fmt.Println("\tlsc <country code>: list city of country")
		fmt.Println("\tcw <city>: view current weather")
	}
}
