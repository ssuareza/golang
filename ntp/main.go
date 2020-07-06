package main

import (
	"fmt"
	"log"
	"time"

	"math"

	"github.com/beevik/ntp"
)

const (
	ntpServer = "pool.ntp.org"
)

func main() {
	response, err := ntp.Query(ntpServer)
	if err != nil {
		log.Fatal(err)
	}

	ntpTime := time.Now().Add(response.ClockOffset)
	fmt.Println("NTP", ntpTime)

	now := time.Now()
	fmt.Println("NOW", now)

	diff := now.Sub(ntpTime)
	fmt.Println("DIFF", math.Round(diff.Seconds()))
}
