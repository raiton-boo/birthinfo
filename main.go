package main

import (
	"fmt"
	"birthinfo/weekday"
)

func main() {
	year := 2007
	month := 1
	day := 19

	inputDate, week := weekday.Find(year, month, day)

	fmt.Printf("%s あなたの産まれた日は %s です\n", inputDate, week)
}