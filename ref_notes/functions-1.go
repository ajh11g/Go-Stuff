// learnin' functions
// function name rules:
//  - name must begin with a letter
//	- functions that begin with a capital letter are exported,
//    can be used outside the current pacakge
//  - use camelCase

package main

import "fmt"

// basic return function
func double(number float64) float64 {
	return number * 2
}

// function with parameters
func repeatLine(line string, times int) {
	for i := 0; i <= times; i++ {
		fmt.Println(line)
	}
}

// simple function
func sayHi() {
	fmt.Println("Hello!")
}

func main() {
	sayHi()
	repeatLine("Hello?", 3)
	fmt.Println(double(4.2))
}
