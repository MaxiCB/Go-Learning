package main

import (
	"fmt"
	"log"
	"aaroncb.com/greetings"
)


func main() {
	log.SetPrefix("grretings: ")
	log.SetFlags(0);
	names := []string{"Aaron", "Lyric", "Turtle"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}