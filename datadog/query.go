// Returns a metric value from Datadog
package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/zorkian/go-datadog-api.v2"
)

func main() {
	apiKey := "YOUR_API_KEY_HERE"
	appKey := "YOUR_APP_KEY_HERE"
	client := datadog.NewClient(apiKey, appKey)
	query := "sum:istio.mesh.request.count{destination_namespace:default,destination_service:myservice,response_code:500}"

	timeFrom := time.Now().Unix()
	time.Sleep(120 * time.Second)
	timeNow := time.Now().Unix()

	series, err := client.QueryMetrics(timeFrom, timeNow, query)
	if err != nil {
		log.Fatal(err)
	}

	var r float64
	for _, s := range series {
		for _, p := range s.Points {
			fmt.Printf("%f %f\n", *p[0], *p[1])
			r += *p[1]
		}
	}
	fmt.Printf("RESULT %f\n", r)
}
