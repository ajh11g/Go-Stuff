// Talking Clock - from reddit.com/r/dailyprogrammer #321 challenge
// Description: turns number-based 24-hour clock into text output
// Refs: /u/QuantumEntanlged's submission - https://bit.ly/2LWCUGN

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	currentTime := time.Now().Format("15:04")
	challengeInput := []string{currentTime, "00:00", "01:30", "12:05", "14:01", "20:29", "21:00"}

	for _, i := range challengeInput {
		input := strings.SplitN(i, ":", 2)
		numHour, _ := strconv.Atoi(input[0])
		numMin, _ := strconv.Atoi(input[1])

		wordHour := hour2word(numHour)
		wordMinute1, wordMinute2 := minute2word(numMin)

		if i == currentTime {
			output := fmt.Sprintf("It's currently %s %s %s %s\n", wordHour, wordMinute1, wordMinute2, amOrPm(numHour))
			fmt.Println(standardizeSpaces(output))
			continue
		}

		output := fmt.Sprintf("It's %s %s %s %s\n", wordHour, wordMinute1, wordMinute2, amOrPm(numHour))
		fmt.Println(standardizeSpaces(output))
	}
}

func minute2word(numMin int) (string, string) {
	if numMin == 0 {
		return "o'clock", ""
	} else if (numMin > 10) && (numMin <= 19) {
		return minutesList3[numMin], ""
	} else if numMin <= 10 {
		return "oh", minutesList1[numMin]
	} else if numMin >= 20 {
		sliceMin := strings.SplitN(strconv.Itoa(numMin), "", 2)
		num2, _ := strconv.Atoi(sliceMin[0])
		num1, _ := strconv.Atoi(sliceMin[1])
		return minutesList2[num2], minutesList1[num1]
	} else {
		fmt.Println("ERROR")
		return "", ""
	}
}

func hour2word(numHour int) string {
	return hoursList[numHour]
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func amOrPm(number int) string {
	if number <= 11 {
		return "am"
	}
	return "pm"
}

var hoursList = map[int]string{
	0:  "twelve",
	1:  "one ",
	2:  "two ",
	3:  "three ",
	4:  "four ",
	5:  "five ",
	6:  "six ",
	7:  "seven ",
	8:  "eight ",
	9:  "nine ",
	10: "ten ",
	11: "eleven ",
	12: "twelve ",
	13: "one ",
	14: "two ",
	15: "three ",
	16: "four ",
	17: "five ",
	18: "six ",
	19: "seven ",
	20: "eight ",
	21: "nine ",
	22: "ten ",
	23: "eleven ",
}

var minutesList1 = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var minutesList2 = map[int]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
}

var minutesList3 = map[int]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}
