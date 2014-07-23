package main

import (
	"fmt"
	"strconv"
)

func main() {
	var output string
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			output = "FizzBuzz"
		} else if i%3 == 0 {
			output = "Fizz"
		} else if i%5 == 0 {
			output = "Buzz"
		} else {
			output = strconv.Itoa(i)
		}

		fmt.Println(output)
	}
}
