package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type MovieA struct {
	Title string `json:"title"`
	Name  string `json:"name"`
}

func main() {
	m1 := Movie{"喜剧之王", 2000}

	jsonStr, err := json.Marshal(m1)

	if err != nil {
		fmt.Println("failed to string to JSON")
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("failed to convert jsonStr to struct")
		return
	}
	fmt.Println(myMovie)

	myMovieA := MovieA{}
	err = json.Unmarshal(jsonStr, &myMovieA)
	if err != nil {
		fmt.Println("failed to convert jsonStr to struct")
		return
	}
	fmt.Println(myMovieA)
}
