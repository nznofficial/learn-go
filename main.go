package main

import (
	"fmt"
	"learngo/calc"
	"learngo/functions"
	"runtime"
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

func double(a int) int {
	return a + a
}

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func printReports(intro, body, outro string) {
	printCostReport(func(a string) int { return len(a) * 2 }, intro)
	printCostReport(func(a string) int { return len(a) * 3 }, body)
	printCostReport(func(a string) int { return len(a) * 4 }, outro)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
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

	fmt.Println(functions.Concat("Elijah", "Vazquez"))

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

	// Switch statements
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Println(os)
	}

	go_x := 1
	switch {
	case go_x == 1:
		fmt.Println("Go 1.1")
		fallthrough
	case go_x == 2:
		fmt.Println("Go 1.2")
	}

	go_x_2 := 10
	if go_x_2 := 5; go_x_2 > 3 {
		fmt.Println(go_x_2)
	}
	fmt.Println(go_x_2)

	for i := 0; i < 3; i++ {
		switch {
		case i == 1:
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		switch {
		case i == 1:
			continue
		}
		fmt.Println(i)
	}

outer:
	for i := 0; i < 3; i++ {
		switch {
		case i == 1:
			break outer
		}
		fmt.Println(i)
	}

	// Chapter 4 - Functions

	fmt.Println(functions.GetMonthlyPrice("basic"))
	fmt.Println(functions.GetMonthlyPrice("premium"))
	fmt.Println(functions.GetMonthlyPrice("enterprise"))
	fmt.Println(functions.GetMonthlyPrice("unknown"))

	cost_per_message := 2
	messagesThisMonth := 50
	messagesLastMonth := 40

	fmt.Println(functions.MonthlyBillIncrease(cost_per_message, messagesLastMonth, messagesThisMonth))

	tier_basic := "basic"
	tier_premium := "premium"
	tier_enterprise := "enterprise"
	tier_unknown := "unknown"

	fmt.Println(functions.GetProductMessage(tier_basic))
	fmt.Println(functions.GetProductMessage(tier_premium))
	fmt.Println(functions.GetProductMessage(tier_enterprise))
	fmt.Println(functions.GetProductMessage(tier_unknown))

	fmt.Println(functions.YearsUntilEvents(15))

	//fmt.Println(functions.Reformat("Hello There", ))

	newX, newY, newZ := conversions(double, 1, 2, 3)
	fmt.Println(newX, newY, newZ)

	newX, newY, newZ = conversions(func(a int) int {
		return a + a
	}, 5, 10, 15)
	fmt.Println(newX, newY, newZ)

	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}
