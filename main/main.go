package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/fizzbuzz/:count", fizzbuzz)
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func fizzbuzz(responseWriter http.ResponseWriter, request *http.Request) {

	var input = 16
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
	} else if request.Method == "GET" {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		respData := fmt.Sprintf("{\"fizzBuzz\": %s}", jsonObj)
		responseWriter.Write([]byte(respData))

	}
}
