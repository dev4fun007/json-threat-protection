package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonStream = `{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}`

	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}

		tokenType := t.(json.Delim)
		fmt.Print(tokenType)

		fmt.Printf("%T : %v", t, t)
		if decoder.More() {
			fmt.Printf(" more ")
		}
		fmt.Printf("\n")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
