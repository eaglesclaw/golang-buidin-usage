package main

import "log"


func main(){


	slice := make([]byte, 10)
	log.Printf("slice: %v, len: %d",slice,len(slice))

	str := "merhaba televole"
	log.Printf("string: %v, len: %d",str,len(str))
}