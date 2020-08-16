package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omtempty"`
	Actors []string
}
func main() {
	var movies = []Movie{
		{Title:"1",Year:1942,Color:false, Actors:[]string{"A","B"}},
		{Title:"2",Year:1942,Color:false, Actors:[]string{"A","B"}},
		{Title:"3",Year:1942,Color:false, Actors:[]string{"A","B"}},
		{Title:"4",Year:1942,Color:false, Actors:[]string{"A","B"}},
	}
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies,"","	")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
