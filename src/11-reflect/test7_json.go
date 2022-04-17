package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	spiderMan := Movie{"spider-man", 2001, 100, []string{"蜘蛛侠", "lisa"}}

	jsonStr, err := json.Marshal(spiderMan)
	if err != nil {
		fmt.Println("json marshal error ", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	ppp := Movie{}
	err = json.Unmarshal(jsonStr, &ppp)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}
	fmt.Printf("%v\n", ppp)
}
