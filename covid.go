package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func statGlobal() {

	statsJson := covidGlobalResult{}
	res, err := http.Get("https://api.covid19api.com/summary")

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "An error occured with the country")
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		defer res.Body.Close()
		return
	}

	err = json.Unmarshal(body, &statsJson)

	fmt.Printf("\nGlobal reports:" +
		"\n\tCases:          %d" +
		"\n\tDeaths:         %d" +
		"\n\tRecoveries:     %d\n\n",
		statsJson.Global.Confirmed,
		statsJson.Global.Deaths,
		statsJson.Global.Recovered)
}

func stat(country string) {
	statsJson := covidCountryResults{}

	res, err := http.Get(fmt.Sprintf("https://api.covid19api.com/total/country/%s", country))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "An error occured with the country %s\n", country)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		defer res.Body.Close()
		_, _ = fmt.Fprintf(os.Stderr, "An error occured with the country %s\n", country)
		return
	}

	err = json.Unmarshal(body, &statsJson)
	if err != nil {
		defer res.Body.Close()
		_, _ = fmt.Fprintf(os.Stderr, "An error occured with the country %s\n", country)
		return
	}
	statsCountryJson := statsJson[len(statsJson) - 1]

	fmt.Printf("\nReports for %s:" +
		"\n\tCases:        %d" +
		"\n\tDeaths:       %d" +
		"\n\tRecoveries:   %d\n",
		statsCountryJson.Country,
		statsCountryJson.Confirmed,
		statsCountryJson.Deaths,
		statsCountryJson.Recovered)
}

func main() {
	args := os.Args

	if len(args) > 1 {
		for _, country := range args[1:] {
			country = strings.ToLower(country)
			country = strings.Replace(country, " ", "-", -1)

			if country == "all" {
				continue
			}
			stat(country)
		}
		fmt.Println()
		os.Exit(0)
	}
	statGlobal()
	os.Exit(0)
}