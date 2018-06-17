// loop learnin'

package main

import (
	"fmt"
)

func main() {

	// for loops
	//  |init statement
	//  |       |condition expression
	//  |       |       |post statement
	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}

	// init and post statements are optional
	someVariable := 1
	for someVariable <= 3 {
		fmt.Println(someVariable)
		someVariable++
	}

	// continue == skip to the next iteartion of a loop
	// break == breaks out of the current loop
	for x := 1; x <= 3; x++ {
		fmt.Println("before continue")
		continue
		// this should never get printed
		fmt.Println("after continue")
	}

	for x := 1; x <= 3; x++ {
		fmt.Println("before break")
		break
		// this should never get printed
		fmt.Println("after break")
	}
	fmt.Println("After break")
}
