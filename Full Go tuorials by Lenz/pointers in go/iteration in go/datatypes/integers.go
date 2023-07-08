// integers are both signed and unsigned
package main

import "fmt"

func main() {
	var u uint8 = 22
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	var l uint8 = 80
	fmt.Println(l, l+1, l*l)

}
