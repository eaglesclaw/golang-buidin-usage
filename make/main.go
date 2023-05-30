package main

import "log"
// 

func main(){
/*
	//v1
	states := make(map[string]string)
	fmt.Println(states)

	states["TEK"] = "Tekirdağ"
	states["IST"] = "Istanbul"
	states["NYC"] = "New York"

	fmt.Println(states)

	ist := states["IST"]
	fmt.Println(ist)

	// silme

	delete(states, "NYC")

	fmt.Println(states)

	//v2
	unbuffered := make(chan int)

	log.Printf("Unbuffered: %v, type %T, len: %d, cap %d",unbuffered,unbuffered,len(unbuffered),cap(unbuffered))

	//v3
	buffered := make(chan int, 10)

	log.Printf("buffered: %v, type %T, len: %d, cap %d",buffered,buffered,len(buffered),cap(buffered))
*/
	//v4
	slice := make([]byte, 7)
	log.Printf("slice: %v, type %T, len: %d, cap %d",slice,slice,len(slice),cap(slice))

	slice2 := make([]byte, 4,11)
	log.Printf("slice2: %v, type %T, len: %d, cap %d",slice2,slice2,len(slice2),cap(slice2))
}

