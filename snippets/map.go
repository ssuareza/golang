package main

import "fmt"

func main() {
	instances := make(map[string]string)

	instances["i-dafadsfad"] = "1.1.1.1"
	instances["i-qerqwerqe"] = "2.2.2.2"
	instances["i-zccvzvvxz"] = "3.3.3.3"

	fmt.Println(instances)

	// check if map entry exists
	if _, ok := instances["i-dafadsfad"]; ok {
		fmt.Println("exists!!!")
	}

	// delete map entry
	delete(instances, "i-zccvzvvxz")

	// print all value
	for _, v := range instances {
		fmt.Println(v)
	}
}
