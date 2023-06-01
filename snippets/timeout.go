package main

import (
	"fmt"
	"time"
)

func something() bool {
	if time.Now().Unix()%2567 == 0 {
		fmt.Println("done")
		return true
	}
	fmt.Println("not done")
	return false
}

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	after := time.After(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			if something() {
				return
			}
		case <-after:
			fmt.Println("timed out")
			return
		}
	}
}
