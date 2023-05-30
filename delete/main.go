package main

import "log"


func main(){
	m:= make(map[string]int)

	log.Println(m)

	m["Person1"] = 38
	m["Person2"] = 83
	m["Person3"] = 64

	log.Println(m)

	delete(m, "Person2")

	log.Println(m)

	m=nil
	// nil hepsini siler
	log.Println(m)
}