package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Represent the structure of the JSON response from the weather API
type weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

// Weather fetches weather information for the specified location using the weather API
func Weather(location string) (string, float64, string) {
	key := "35bd76d57c7f4652add204524240102"

	URL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", key, location)

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	//Perform the API request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	//Read the response body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	//Unmarshal the JSON response into the weather struct
	var responseObject weather
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject.Location.Name, responseObject.Current.TempC, responseObject.Current.Condition.Text
}
