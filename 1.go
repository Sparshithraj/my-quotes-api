package main

import "fmt"

func main() {
	adder(add)
	adder1(add)
}
func adder(add func(int, int) int) {
	val := add(3, 5)
	fmt.Println(val)
}

type myFunctionDataType func (int, int) int

func adder1(myFunctionDataType) {
	val := add(3, 5)
	fmt.Println(val)
}

func add(i int, j int) int {
	return i + j
}