package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://montanaflynn.github.io/covid-19/data/current.json")
	fmt.Println(resp)
	var in map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&in)
	fmt.Println(in)
}
