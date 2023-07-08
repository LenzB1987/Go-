package main

import "fmt"

func main() {
	//var f=fr
	var freezing_point_in_F = 32.0
	var boiling_point_in_F = 212.0
	fmt.Printf("%gF= %gC\n", freezing_point_in_F, faren_to_celcz(freezing_point_in_F))
	fmt.Printf("%gF=%gC\n", boiling_point_in_F, faren_to_celc(boiling_point_in_F))
}
func faren_to_celcz(float64) float64 {
	return (32.0 - 32.0*5/9)
}
func faren_to_celc(float64) float64 {
	return (212.0 - 212.0*5/9)
}
