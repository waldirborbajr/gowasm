package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// var htmlString = `<h4>Hello, I'm an HTML snippet from Go!</h4>`

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

// func GetHtml() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		return htmlString
// 	})
// }

func main() {

	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)

	// ch := make(chan struct{}, 0)
	// fmt.Printf("Hello Web Assembly from Go!\n")

	// js.Global().Set("getHtml", GetHtml())
	// <-ch
}
