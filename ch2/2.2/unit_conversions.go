package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"paulghaddad/gobook/ch2/distconv"
	"paulghaddad/gobook/ch2/tempconv"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		stdin, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from STDIN: %v\n", err)
			os.Exit(1)
		}
		convertNums(strings.Split(strings.TrimSpace(string(stdin)), "\n"))
	} else {
		convertNums(os.Args[1:])
	}
}

func convertNums(nums []string) {
	for _, num := range nums {
		num, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid argument: %v\n", err)
			os.Exit(1)
		}

		fahrenheit := tempconv.Fahrenheit(num)
		celsius := tempconv.Celsius(num)
		fmt.Printf("%gF = %s, %gC = %s\n", num, tempconv.FToC(fahrenheit), num, tempconv.CToF(celsius))

		feet := distconv.Feet(num)
		meters := distconv.Meter(num)
		fmt.Printf("%gF = %s, %gM = %s\n", num, distconv.FToM(feet), num, distconv.MToF(meters))
	}
}
