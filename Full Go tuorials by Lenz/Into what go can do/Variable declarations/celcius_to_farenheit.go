package main

import "fmt"

//boiling point of water
const boiling_point = 212.0

func main() {
	var f = boiling_point
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
	//prints 100 c 0r 212  f

}
