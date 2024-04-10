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

func (b *Burger) WithCheese() *Burger {
	b.Price += 1.50
	b.Ingredients = append(b.Ingredients, "cheese")

	// NOTE: return 'b' is the beauty of builder (allowing method call concatenation on the return)
	return b
}

func (b *Burger) WithTomato() *Burger {
	b.Price += 0.30
	b.Ingredients = append(b.Ingredients, "tomato")

	return b
}

func (b *Burger) DoubleBurger() *Burger {
	b.Price += 5.00
	b.Ingredients = append(b.Ingredients, "meat")

	return b
}

func (b *Burger) WithKetchup() *Burger {
	b.Price += 0.20
	b.Ingredients = append(b.Ingredients, "ketchup")

	return b
}
