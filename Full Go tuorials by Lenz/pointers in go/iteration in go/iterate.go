package main

import "fmt"

func main() {
	x := "hello!"
	//iterating through hello
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
			//prints HELLO
		}
	}
}
