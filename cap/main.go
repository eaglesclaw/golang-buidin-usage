package main

import (
	"log"
)

func main(){

	// v1
	slice := make([]byte, 0, 5)
	log.Printf("slice: %d", cap(slice))

	// v2

	channel := make(chan int, 10)
	log.Printf("channel: %d", cap(channel))
	
	//v3
	var pointer *[15]byte
	log.Printf("pointer: %d == %d", cap(pointer), len(pointer))

}