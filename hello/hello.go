package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

// message, err := greetings.Hello("Khanh")

names := []string{"Khanh", "Le"}
 messages, errs := greetings.Hellos(names)

//  if err != nil || errs != nil{
// 	log.Fatal(err)
//  }

 //fmt.Println(message)

 if errs != nil {
	log.Fatal(errs)
 }
 fmt.Println(messages)
}