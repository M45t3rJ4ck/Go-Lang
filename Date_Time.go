package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), now.Minute(), now.Second())

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	fmt.Printf("Current time in Tokyo: %s\n", now.In(loc))

	birthDate := time.Date(1983, time.May, 27, 14, 0, 0, 0, time.Local)
	fmt.Printf("Birth date: %s\n", birthDate.Format("2006-01-02 15:04:05"))
	diff := now.Sub(birthDate)
	fmt.Printf("Time since birth: %d days\n", int(diff.Hours()/24))
	fmt.Printf("Hours since birth: %d hours\n", int(diff.Hours()))
}
