package feature

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Country string   `json:"country"`
	Cities  []string `json:"cities"`
}

func ListCity(country string){
	data, err := os.ReadFile("city.json")
	if err != nil{
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}

	var countries []Data

	if err := json.Unmarshal(data, &countries); err != nil{
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}

	for _, data := range countries{
		if data.Country == country{
			for _, city := range data.Cities{
				fmt.Println(city)
			}
		}
	}
}