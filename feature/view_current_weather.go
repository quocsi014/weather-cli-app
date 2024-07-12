package feature

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type BodyRes struct {
	Weather []Weather `json:"weather"`
	Deg Deg `json:"main"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Deg struct{
	Degree float32 `json:"temp"`
}

func (d *Deg)ToCelsius(){
	d.Degree = d.Degree - 272.15
}


func ViewCurrentWeather(city string) {
	res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", strings.ReplaceAll(city, "_", "%20"), os.Getenv("API_KEY")))
	if err != nil {
		fmt.Println("Something went wrong with server, try again pls")
		return
	}
	if res.StatusCode == 401 {
		fmt.Println("Your api key has a problem, please check your api key again")
		fmt.Println("You can change the code with command: ckey <key>")
		println("If you dont have a key, you can get it at https://openweathermap.org")
		return
	}
	if res.StatusCode == 404{
		fmt.Println("No city found")
		fmt.Println("Try with another city, or you can see cities of your country by command:")
		fmt.Println("lsc <country_code>")
		return
	}
	defer res.Body.Close()

	var data []byte

	if _, err := res.Body.Read(data); err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}

	var bodyRes BodyRes

	if err := json.NewDecoder(res.Body).Decode(&bodyRes); err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
	}

	bodyRes.Deg.ToCelsius()

	fmt.Printf("Status: %s\n", bodyRes.Weather[0].Main)
	fmt.Printf("Description: %s\n", bodyRes.Weather[0].Description)
	var c *color.Color = color.New(color.FgWhite)
	if bodyRes.Deg.Degree < 15{
		c.Add(color.BgBlue)
	} else if bodyRes.Deg.Degree >25{
		c.Add(color.BgRed)
	} else{
		c.Add(color.BgHiGreen)
	}

	c.Println(fmt.Sprintf("Temp: %.2f°C", bodyRes.Deg.Degree))

}
