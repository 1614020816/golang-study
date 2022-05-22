package main

import "fmt"

func main() {
	var a string

	// pair<statictype: string , value: "success">
	a = "success"

	// pair<statictype: string , value: "success">
	var allType interface{}
	allType = a

	str, _ := allType.(string)

	fmt.Println(str)
	fmt.Printf("allType type is %T\n", allType)

}