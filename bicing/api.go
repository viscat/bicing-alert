package bicing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const API_URI = "http://wservice.viabicing.cat/v2"

func GetStationsStatus() (Status, error) {
	resp, err := http.Get(fmt.Sprintf("%v/stations", API_URI))
	defer resp.Body.Close()
	if err != nil {
		return Status{}, fmt.Errorf("Error retrieving stations status: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return Status{}, fmt.Errorf("Error parsing response: %v", err)
	}
os.Getwd()
	var status Status
	err = json.Unmarshal(body, &status)

	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling bicing api response: %v", err))
	}

	return status, nil
}
