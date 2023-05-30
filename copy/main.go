package main

import "fmt"


func main(){
	/*
	//v1
	var s = make([]int,5)
	n := copy(s, []int{1,2,3,4})

	// 3 kayıt diğer tarafa aktarıldı
	fmt.Println(n)

	// kapasitesi kadar elemanı aldı
	fmt.Println(s)
	*/


	/*
	s:= []int{0,1,2,3}
	n:= copy(s, s[1:])

	fmt.Println(n)

	fmt.Println(s)
	*/

	var x = make([]byte,5)
	mystr:= "Hello mars!"
	copy(x, mystr)
	fmt.Println(mystr)

	fmt.Println(string(x))
}