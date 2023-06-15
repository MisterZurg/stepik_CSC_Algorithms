package main

import "fmt"

// MoneyChange() implements a simple greedy algorithm used by cashiers all over the world.
func MoneyChange(money int) int {
	// denominations 1, 5, 10
	if money%10 == 0 {
		return money / 10
	}

	change := 0
	change += money / 10
	// Number of 5 cent coins
	money %= 10
	if money%5 == 0 {
		return change + money/5
	}
	change += money / 5
	money %= 5
	return change + money
}

func main() {
	// put your code here
	var money int
	fmt.Scan(&money)
	fmt.Println(MoneyChange(money))
}
