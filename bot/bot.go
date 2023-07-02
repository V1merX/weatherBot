package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type weatherBot struct {
	appVersion float32
	appCity    string
	apiKey     string
}

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

func Start(apiKey string, appVersion float32) {
	bot := new(weatherBot)

	if apiKey == "<API-KEY>" {
		log.Fatal("Вы не ввели api ключ. Определение погоды невозможно")
	}

	bot.apiKey = apiKey
	bot.appVersion = appVersion

	bot.scanCity()

	bot.sendData().printResult()
}

func (wD weatherData) printResult() {
	fmt.Println("Ваши данные о погоде из " + wD.Name + ":")
	for _, weather := range wD.Weather {
		fmt.Println("-", weather.Main)
		fmt.Println("-", weather.Description)
	}
	fmt.Println("-", wD.Wind.Speed)
}

func (wB *weatherBot) scanCity() {
	fmt.Print("Город: ")
	_, err := fmt.Scan(&wB.appCity)
	if err != nil {
		log.Fatal(err)
	}
}

func (wB weatherBot) sendData() weatherData {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + wB.appCity + "&appid=" + wB.apiKey)
	if err != nil {
		log.Fatal(err)
	}

	return wB.getData(resp)
}

func (wB weatherBot) getData(resp *http.Response) weatherData {
	defer resp.Body.Close()

	var weatherData weatherData
	err := json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		log.Fatalln(err)
	}

	return weatherData
}
