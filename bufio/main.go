package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// print only words
	input := "foo   bar      baz"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
