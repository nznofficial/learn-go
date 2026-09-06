package main

import (
	"fmt"
	"learngo/calc"
	"unicode/utf8"
)

func billingCost(plan string) float64 {
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterpise":
		return 50.0
	default:
		return 0.0
	}
}

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

	fmt.Println(calc.Add(3, 8))

	// Chapter 3 - Conditionals

	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Sending message:", messageLen, "Max Message:", maxMessageLen)

	if messageLen > maxMessageLen {
		fmt.Println("Message too long")
	} else {
		fmt.Println("Message Sent")
	}

	plan := "basic"
	fmt.Println("Plan:", plan, billingCost(plan))
	plan = "pro"
	fmt.Println("Plan:", plan, billingCost(plan))
	plan = "enterpise"
	fmt.Println("Plan:", plan, billingCost(plan))
	plan = "unknown"
	fmt.Println("Plan:", plan, billingCost(plan))

	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	finalCost = bulkMessageCost

	if isPremiumUser == true {
		finalCost = finalCost - (finalCost * discountRate)
	}

	if finalCost > accountBalance {
		fmt.Println(insufficientFundMessage)
	} else {
		fmt.Println(purchaseSuccessMessage)
		accountBalance = accountBalance - finalCost
	}

	fmt.Println("Account balance:", accountBalance)

	// Chapter 4 - Functions

}
