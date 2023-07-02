# weatherBot
<b>weatherBot</b> is a simple weather identifier for the desired region, written in Go.<br>
There is only a small snippet of interaction with the API in the program. 

<hr>

## Installation
1. Register on [OpenWeatherMap](https://openweathermap.org) and get your first API key
1. Copy the key and enter it in <b>main.go</b> apiKey
2. Start the programme and enter your city.

<hr>

## Adding new functionality
Open the <b>bot/bot.go</b> file and consider the following structure:
```go
type weatherData struct {
	Name string `json:"name"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}
```
On the website page (api integration), find all the information you want to output and add it to this code.<br>
Don't forget to output it in the printResult function

```go
func (wD weatherData) printResult() {
	fmt.Println("Ваши данные о погоде из " + wD.Name + ":")
	for _, weather := range wD.Weather {
		fmt.Println("-", weather.Main)
		fmt.Println("-", weather.Description)
	}
	fmt.Println("-", wD.Wind.Speed)
}
```
