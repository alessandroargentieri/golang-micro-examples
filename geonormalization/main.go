package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type NominatimResult struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

func normalizeCity(cityName string) (string, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Add("q", cityName)
	params.Add("format", "json")
	params.Add("limit", "1")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var results []NominatimResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", fmt.Errorf("city not found")
	}

	return results[0].DisplayName, nil
}

func normalizeAddress(country, postalcode, cityName, street, housenumber string) (string, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	//params.Add("q", cityName)
	params.Add("city", cityName)
	params.Add("street", street)
	params.Add("housenumber", housenumber)
	params.Add("postalcode", postalcode)
	params.Add("country", country)
	params.Add("format", "json")
	params.Add("limit", "1")

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var results []NominatimResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", fmt.Errorf("city not found")
	}

	return results[0].DisplayName, nil
}

func main() {
	fmt.Println(normalizeCity("S.Martino Siccomario"))
	fmt.Println(normalizeCity("Francavilla Fontana 72021 via Occhibianchi 36"))
	fmt.Println(normalizeAddress("italia", "72021", "Francavilla Fontana", "via Occhibianchi", "36"))
	fmt.Println(normalizeAddress("", "", "Francavilla Fontana", "", ""))
	fmt.Println(normalizeAddress("", "", "Francavilla Fontana", "via nonesiste", ""))
	fmt.Println(normalizeAddress("", "", "S.Angelo Lodigiano", "", ""))
	/*
	   output:

	   San Martino Siccomario, Pavia, Lombardia, 27028, Italia <nil>
	   Via Occhibianchi, Francavilla Fontana, Brindisi, Puglia, 72021, Italia <nil>
	   Via Occhibianchi, Francavilla Fontana, Brindisi, Puglia, 72021, Italia <nil>
	   Francavilla Fontana, Brindisi, Puglia, 72021, Italia <nil>
	   city not found
	   Sant'Angelo Lodigiano, Lodi, Lombardia, 26866, Italia <nil>
	*/
}
