package main

import (
	"encoding/json"
	"fmt"
)

type Moive struct {
	Title   string  `json:"title"`
	Year    int     `json:"year"`
	Price   int     `json:"price"`
	Actors  []string `json:"actors"`
}

func main() {
	moive := Moive{"功夫", 2001, 20, []string{"星爷", "张柏芝"}}

	// 解析的过程  struct ========> json
	jsonStr, err := json.Marshal(moive)

	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 编码成结构体的过程  json ========> struct
	// jsonStr = {"title":"功夫","year":2001,"price":20,"actors":["星爷","张柏芝"]}

	myMoive := Moive{}
	err = json.Unmarshal(jsonStr, &myMoive)

	if err != nil {
		fmt.Println("json unmarshal error: ", err)
		return
	}

	fmt.Printf("%v\n", myMoive)
}
