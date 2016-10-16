package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	geoBlog struct{}
	// GeoIP information about where you are
	GeoIP struct {
		IP          string  `json:"ip"`
		CountryCode string  `json:"country_code"`
		CountryName string  `json:"country_name"`
		RegionCode  string  `json:"region_code"`
		RegionName  string  `json:"region_name"`
		City        string  `json:"city"`
		ZipCode     string  `json:"zip_code"`
		TimeZone    string  `json:"time_zone"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	}
)

func (b geoBlog) collect(ch chan string) {
	client := http.Client{}
	r, err := client.Get("http://freegeoip.net/json")
	if err != nil {
		fmt.Printf("Failed geo ip: %v", err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf("to read body: %v", err)
		return
	}
	ip := GeoIP{}
	if err = json.Unmarshal(body, &ip); err != nil {
		fmt.Printf("Unable to parse json: %v", err)
		return
	}
	ch <- fmt.Sprintf("Located at %s (%v : %v)\n", ip.City, ip.Latitude, ip.Longitude)
}
