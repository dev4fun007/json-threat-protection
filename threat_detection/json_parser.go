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

	const jsonStream = `[{"Message": "Hello", "Number": 1.234}]`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	//Each entry will store count of array elements - index being array level
	//0 -> count of 1st level array elements, 1 -> count of 2nd level array elements
	lenCountSlice := make([]int, 0)
	//Indicates the array level
	bracesCount := 0
	//var stack []string

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		tValue := fmt.Sprintf("%v", t)
		if tValue == "[" {
			//Array Start
			//Add to stack braces in the format - [1, [2, [3
			//Each nested array bracket will increment the count
			//stack = append(stack, tValue + string(bracesCount))
			bracesCount++
		} else if tValue == "]" {
			//Array End
			//Pop the top entry from stack
			//...[3, [2, [1 - This is how entry will be popped
			//n := len(stack) -1
			//stack = stack[:n]
			bracesCount--
		}

		if bracesCount > 0 {
			//The incoming tokens are elements of an array
			//The current bracesCount value indicates the current array level in the json
			//Increment the count value at bracesCount index
			if len(lenCountSlice) > bracesCount {
				//Already have a value for this bracesCount index, as length of the slice is greater
				lenCountSlice[bracesCount] = lenCountSlice[bracesCount] + 1
			} else {
				//Length of slice is less - slice does not have an entry at bracesCount index
				//Append at this index with count as 1
				lenCountSlice = append(lenCountSlice, 1)
			}
		}
	}
}

func validateDepth() {
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
	}

	fmt.Println("Max Depth", maxDepth)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
