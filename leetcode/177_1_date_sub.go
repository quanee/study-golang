package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	fmt.Println(daysBetweenDates(date1, date2))
}

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)
	return int(d1.Sub(d2).Hours() / 24)
}
