package main

import (
	"fmt"
)

func main() {
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

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("permissions:", permissions)
	fmt.Println("costPerSMS:", costPerSMS)
	fmt.Println(messageStart, age, messageEnd)
	fmt.Println("Plan:", basicPlan)
}
