package main

import (
	"fmt"
)

var (
	water       = 400
	milk        = 540
	coffeeBeans = 120
	cups        = 9
	money       = 550
)

func main() {
	var (
		stopMachine bool
		action      string
	)

	for !stopMachine {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		_, _ = fmt.Scanln(&action)

		switch action {
		case "buy":
			buy()
		case "fill":
			fill()
		case "take":
			takeMoney()
		case "remaining":
			dispMaterials()
		case "exit":
			stopMachine = true
			return
		}
		fmt.Println()
	}
}

// BUY ACTION
func buy() {
	var coffeeType int
	var op string
	fmt.Println("\nWhat do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	_, _ = fmt.Scanf("%s", &op)
	switch op {
	case "1":
		coffeeType = 1
	case "2":
		coffeeType = 2
	case "3":
		coffeeType = 3
	}
	// verify the possibility of do coffee
	if op != "back" {
		makeCoffee(coffeeType)
	}
}

func makeCoffee(t int) {
	notEnough := [...]string{"water", "milk", "coffee beans", "cups"}
	var log int
	//TODO: SEPARATE materials AND ADD EACH COFFEE TYPE TO AN INDEPENDENT ARRAY AND JOIN IT
	// 1 - espresso, 2 - latte, 3 - cappuccino
	materials := [3][4]int{{250, 0, 16, 4},
		{350, 75, 20, 7},
		{200, 100, 12, 6}}
	t -= 1

	log = tryDo(materials[t][0], materials[t][1], materials[t][2])

	if log == len(notEnough) {
		// do(t)
		do(materials[t][0], materials[t][1], materials[t][2], materials[t][3])
		fmt.Println("I have enough resources, making you a coffee!")
	} else {
		fmt.Printf("Sorry, not enough %s!\n", notEnough[log])
	}
}

func do(w, m, c, cash int) {
	material := [...]*int{&water, &milk, &coffeeBeans, &cups}
	used := [...]int{w, m, c, 1}

	for i := range material {
		*material[i] -= used[i]
	}
	money += cash
}

func tryDo(w, m, c int) int {
	dispWater, dispMilk, dispCoffeeBeans, dispCups := water, milk, coffeeBeans, cups
	material := [...]int{dispWater, dispMilk, dispCoffeeBeans, dispCups}
	used := [...]int{w, m, c, 1}

	for i := range material {
		material[i] -= used[i]
		if material[i] < 0 {
			return i
		}
	}
	return len(material)
}

// FILL ACTION
func fill() {
	waterAdd, milkAdd, coffeeBeansAdd, cupsAdd := 0, 0, 0, 0
	// take materials
	// waterAdd := ask("Write how many ml of water you want to add:")
	fmt.Println("\nWrite how many ml of water you want to add: ")
	fmt.Scanln(&waterAdd)
	fmt.Println("Write how many ml of milk you want to add: ")
	fmt.Scanln(&milkAdd)
	fmt.Println("Write how many grams of coffee beans you want to add: ")
	fmt.Scanln(&coffeeBeansAdd)
	fmt.Println("Write how many disposable cups you want to add: ")
	fmt.Scanln(&cupsAdd)
	// save values
	water += waterAdd
	milk += milkAdd
	coffeeBeans += coffeeBeansAdd
	cups += cupsAdd
}

// TAKE ACTION
func takeMoney() { // prints money
	fmt.Printf("I gave you $%d\n", money)
	money = 0
}

// REMAINING ACTION
func dispMaterials() { // prints disposable materials
	fmt.Println("\nThe coffee machine has:")
	fmt.Printf("%d ml of water\n", water)
	fmt.Printf("%d ml of milk\n", milk)
	fmt.Printf("%d g of coffee beans\n", coffeeBeans)
	fmt.Printf("%d disposable cups\n", cups)
	fmt.Printf("$%d of money\n", money)
}
