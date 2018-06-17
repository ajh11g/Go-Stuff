// formatting stuff

package main

import (
	"fmt"
)

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height

	fmt.Println(area/10.0, "liters needed")
	// use printf "print, with formatting"
	fmt.Printf("%0.2f liters needed\n", area/10.0)
	// or use Sprintf, which returns a string
	resultString := fmt.Sprintf("%0.2f liters needed\n", area/10.0)
	fmt.Printf(resultString)

	// formatting verbs examples
	// - %f=floating point number
	// - %d=decimal integer
	// - %s=string
	// - %t=boolean
	// - %v=any value
	//              |verb   |verb
	fmt.Printf("The %s cost %d cents each.\n", "gumballs", 25)
	fmt.Printf("Or %f if you buy 6 of them!\n", 0.22*6)
	fmt.Printf("Types: %T %T %T\n\n", 1.2, "\t", true)

	// fomratting value widths
	//          |min width of 12 chars
	//          |      |no min width
	fmt.Printf("%12s | %s\n", "Product", "Cost in cents")
	fmt.Println("-----------------------------")
	//                 |min width of 2
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n\n", "Paper clips", 5)

	// formatting fractional number widths
	// start of format sequence
	// |min width of entire number
	// || width after decimal point
	// || |format sequence type
	// %4.3f
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456)
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456)
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456)
	fmt.Printf("%%.1f: %.1f\n", 12.3456)
}
