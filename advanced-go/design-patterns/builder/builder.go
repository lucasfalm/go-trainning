package builder

type Burger struct {
	Price       float64
	Ingredients []string
}

func OrderBurger() *Burger {
	return &Burger{
		Price:       8.00,
		Ingredients: []string{"bread", "bread", "meat"},
	}
}

func (x *Burger) WithCheese() *Burger {
	x.Price += 1.50
	x.Ingredients = append(x.Ingredients, "cheese")

	// NOTE: return 'x' is the beauty of builder (allowing method call concatenation on the return)
	return x
}

func (x *Burger) WithTomato() *Burger {
	x.Price += 0.30
	x.Ingredients = append(x.Ingredients, "tomato")

	return x
}

func (x *Burger) DoubleBurger() *Burger {
	x.Price += 5.00
	x.Ingredients = append(x.Ingredients, "meat")

	return x
}

func (x *Burger) WithKetchup() *Burger {
	x.Price += 0.20
	x.Ingredients = append(x.Ingredients, "ketchup")

	return x
}
