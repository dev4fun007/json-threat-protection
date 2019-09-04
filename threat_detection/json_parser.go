package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"strings"
)

func main() {
	const jsonStream = `{"Message": {"Hello" : {"Hello" : "Hello"}}, "Array": [1, 2, 3], "Null": null, "Number": 1.234}`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))
	var depthStack []string
	maxDepth := 0
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}

		token := fmt.Sprintf("%v", t)
		if token == "{" || token == "[" {
			//push to stack
			depthStack = append(depthStack, token)
			maxDepth = int(math.Max(float64(maxDepth), float64(len(depthStack))))
		} else if token == "}" || token == "]" {
			//pop from stack
			n := len(depthStack) - 1
			depthStack = depthStack[:n]
		}

		/*fmt.Printf("%T : %v", t, t)
		if decoder.More() {
			fmt.Printf(" more ")
		}
		fmt.Printf("\n")*/
	}

	fmt.Println("Max Depth", maxDepth)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
