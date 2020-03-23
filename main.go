package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Do you want to convert a decimal number to binary?")
	fmt.Println("Type the number you want to convert to binary: ")
	var decimalNumber int
	fmt.Scan(&decimalNumber)
	fmt.Printf("Here is the number in decimal and the binary equivalent below %d\n", decimalNumber)
	binaryNumber := int64(decimalNumber)
	fmt.Println(strconv.FormatInt(binaryNumber, 2))
}
