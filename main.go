package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Chapter 2 - Constants and Formatting
	var username string
	username = "chomp_bot"

	var isAdmin bool
	isAdmin = true

	var permissions int
	permissions = 0x1F

	var costPerSMS float64
	costPerSMS = 0.05

	messageStart := "Happy Birthday! You are now"
	age := 40
	messageEnd := "years old!"

	const basicPlan = "basic"

	const secondsInHour = 3600

	const name = "Saul Goodman"
	const openRate = 30.54
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent \n", name, openRate)

	const bear = "🐻"
	fmt.Printf("constant 'name' rune length: %d\n", utf8.RuneCountInString(bear))
	fmt.Println(bear)

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("permissions:", permissions)
	fmt.Println("costPerSMS:", costPerSMS)
	fmt.Println(messageStart, age, messageEnd)
	fmt.Println("Plan:", basicPlan)
	fmt.Println("secondsInHour:", secondsInHour)
	fmt.Print(msg)

}
