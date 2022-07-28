package main

import "fmt"

func div(a float32, b float32) {

	var sonuc float32

	sonuc = a / b
	fmt.Print(sonuc)
}

func main() {

	var x float32
	x = 22.0
	var y float32
	y = 7.0
	div(x, y)
}
