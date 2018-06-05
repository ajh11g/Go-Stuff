// basic grade calculator

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var status string

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	grade, _ := strconv.ParseFloat(input, 64)

	if grade >= 60 {
		status = "passing"

	} else {
		status = "failing"
	}

	fmt.Println(status)
}
