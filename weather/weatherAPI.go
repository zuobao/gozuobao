package weather
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"log"
)

type HeWeather struct {
	Apikey, Url	string
}

type HeWeatherResult struct {
	Basic struct {
		City		string	`json:"city"`
		CityCode	string  `json:"id"`
		Country		string  `json:"cnty"`
		Lat			string 	`json:"lat"`
		Lon			string 	`json:"lon"`
	}
	Status	string    `json:"status"`
	Now	struct {
		Cond struct {
			Code string `json:"code"`
			Txt	 string `json:"txt"`
		}
	}
	DailyForecast []struct {
		Date string `json:"date"`
		Cond	struct {
			CodeD	string `json:"code_d"`
			CodeN	string `json:"code_n"`
			TxtD	string `json:"txt_d"`
			TxtN	string `json:"txt_n"`
		}
		Tmp	struct {
			Max	string    `json:"max"`
			Min	string    `json:"min"`
		}
	} `json:"daily_forecast"`
}

type Result struct {
	WeatherResult []*HeWeatherResult `json:"HeWeather data service 3.0"`

}

func (me *HeWeather) GetHeWeather(ip string) (*HeWeatherResult, error) {

	requestUrl := me.Url + "cityip=" + ip + "&key=" + me.Apikey
	log.Println("IP：", ip)
//	log.Println("请求地址：", requestUrl)

	resp, err := http.DefaultClient.Get(requestUrl)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.New("网络错误: " + err.Error())
	}

	weatherResult := Result{}

	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf,&weatherResult)
	if err != nil {
		return nil, err
	}

	if weatherResult.WeatherResult[0].Status == "ok" {
		return weatherResult.WeatherResult[0], err

	}else{
		return nil, errors.New(weatherResult.WeatherResult[0].Status)
	}
	return nil, nil
}
