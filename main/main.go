package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter integer: ")
	var input int
	_, _ = fmt.Scanf("%d", &input)

	for i := 1; i <= input; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(input int) {

	fizz := "fizz"
	buzz := "buzz"
	var output []string
	for i := 1; i <= input; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, fizz+buzz)
			fmt.Println(i, fizz+buzz)
		} else if i%3 == 0 {
			output = append(output, fizz)
			fmt.Println(i, fizz)
		} else if i%5 == 0 {
			output = append(output, buzz)
			fmt.Println(i, buzz)
		} else {
			output = append(output, strconv.Itoa(i))
			fmt.Println(i)
		}
	}
	jsonObj, err := json.Marshal(output)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Printf("{\"fizzBuzz\": %s}", jsonObj)
	}
}
