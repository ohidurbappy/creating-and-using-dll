package main

import "C"

import (
	_ "fmt"
)

// func init() {
// 	fmt.Println("init()")
// }

//export Square
func Square(number int) int {
	return number * number
}

func main() {

}
