package main

import (
	"fmt"

	"github.com/lucasfalm/go-training/advanced-go/design-patterns/builder"
	"github.com/lucasfalm/go-training/advanced-go/design-patterns/decorator"
)

func main() {
	// showDecorator()
	showBuilder()
}

func showBuilder() {
	burger := builder.OrderBurger()

	// NOTE: the builder is the ability to build complex objects by
	// 			 calling methods that will handle the complexity behind
	// 			 and they allow concatenation (calling one method on top of the other's return)
	burger.WithCheese().WithTomato().WithKetchup().DoubleBurger()

	fmt.Printf("burger: %v - total of: %v\n", burger.Ingredients, burger.Price)
}

func showDecorator() {
	var (
		// NOTE: food is an interface, which means it does not
		// 			 matter the concrete type, it just need to satisfy the food interface
		food decorator.Food
	)

	food = &decorator.XBurger{}

	// NOTE: x-burger delicious 13.99
	fmt.Println(food.Description(), food.Price())

	// NOTE: adding more behaviors to the x-burger methods
	food = &decorator.Coke{Food: food}
	food = &decorator.Sweet{Food: food}

	// NOTE: x-burger delicious, with coke, with sweet 23.13
	fmt.Println(food.Description(), food.Price())

	// ----- x ----- x ----- x //

	var (
		foodTwo decorator.Food
	)

	foodTwo = &decorator.Coke{}
	foodTwo = &decorator.Sweet{Food: foodTwo}

	fmt.Println(foodTwo.Description(), foodTwo.Price())
}
