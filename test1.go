package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const form = "2006-01-02 15:04:05"

func isLeapYear(y int) {
	// fmt.Println(y)
	if y%4 == 0 && y%100 != 0 || y%400 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not a Leap Year")
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// var line string
	// fmt.Scanln(&line)

	t, _ := time.Parse(form, line)
	// fmt.Println(t)
	year, _, _ := t.Date()
	isLeapYear(year)
}
