package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reformatDate("20th Oct 2052"))
	fmt.Println(reformatDate("6th Jun 1933"))
	fmt.Println(reformatDate("26th May 1960"))
}

func reformatDate(date string) string {
	months := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	dates := strings.Split(date, " ")

	dayStr := dates[0]
	dayStr = dayStr[:len(dayStr)-2]
	if len(dayStr) < 2 {
		dayStr = "0" + dayStr
	}

	monthStr := dates[1]
	monthStr = months[monthStr]

	yearStr := dates[2]

	return yearStr + "-" + monthStr + "-" + dayStr
}
