package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/lucasfalm/go-training/advanced-go/design-patterns/builder"
	"github.com/lucasfalm/go-training/advanced-go/design-patterns/decorator"
	"github.com/lucasfalm/go-training/advanced-go/design-patterns/observer"
)

func main() {
	showObserver()

	// showDecorator()

	// showBuilder()

}

func showBuilder() {
	burger := builder.OrderBurger()

	// NOTE: the builder is the ability to build complex objects by
	// calling methods that will handle the complexity behind
	// and they allow concatenation (calling one method on top of the other's return)
	burger.WithCheese().WithTomato().WithKetchup().DoubleBurger()

	fmt.Printf("burger: %v - total of: %v\n", burger.Ingredients, burger.Price)
}

func showDecorator() {
	var (
		// NOTE: food is an interface, which means it does not
		// matter the concrete type, it just need to satisfy the food interface
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

var (
	MyFuncSubject = observer.EventSubject{Observers: sync.Map{}}
)

func showObserver() {
	var (
		observerOne = observer.EventObserver{ID: 1, Time: time.Now()}
		observerTwo = observer.EventObserver{ID: 2, Time: time.Now()}
	)

	MyFuncSubject.AddListener(&observerOne)
	MyFuncSubject.AddListener(&observerTwo)

	one, two, three := something(1, 2, 3)
	fmt.Println(one, two, three)

	switch one {
	case 1002:
		fmt.Println("1002")
		fallthrough
	default:
		fmt.Println("will also pass here")
	}

	c := make(chan bool)

	go usingChannels(c)

	fmt.Println("before goroutine")

	c <- true

	// NOTE: the same goroutine cannot send and receive, so it will
	// block here until another goroutine sends a signal back
	<-c
}

func something(a, _b, _c int) (one, two, three int) {
	event := observer.Event{Data: observer.DataSchema{Title: "hey you", Body: "wanna have some fun?"}}
	MyFuncSubject.Notify(event)

	one = a + 1
	changingValues(&one, 1000)
	two = one + 2
	three = *(multiplier())
	return
}

func multiplier() (number *int) {
	i := rand.Intn(100*int(time.Millisecond)) / 100000
	number = &i
	return
}

func changingValues(value *int, newValue int) {
	*value = newValue + *value
	return
}

func usingChannels(c chan bool) {
	if <-c { // it get the value from the channel
		time.Sleep(3 * time.Second)
		fmt.Println("it was true")
	}

	c <- true // sends a signal back and releases the showObserver function
}
