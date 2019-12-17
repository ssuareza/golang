package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/zorkian/go-datadog-api.v2"
)

var apiKey = "YOUR API_KEY HERE"
var appKey = "YOUR APP_KEY HERE"

func main() {
	client := datadog.NewClient(apiKey, appKey)

	// parsing json
	j, err := ioutil.ReadFile("monitor.json")
	if err != nil {
		log.Fatal("Not able to read file")
	}

	var monitor *datadog.Monitor
	if err := json.Unmarshal(j, &monitor); err != nil {
		log.Fatal("Error decoding json")
	}

	// check if exists
	current, err := client.GetMonitorsByName(monitor.GetName())
	if err != nil {
		log.Fatal(err)
	}

	// update
	if len(current) > 0 {
		current[0].SetMessage(monitor.GetMessage())
		current[0].SetOptions(monitor.GetOptions())
		current[0].SetQuery(monitor.GetQuery())
		if err := client.UpdateMonitor(&current[0]); err != nil {
			log.Fatal(err)
		}
	}

	// create
	_, err = client.CreateMonitor(monitor)
	if err != nil {
		log.Fatal(err)
	}
}
