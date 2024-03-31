package decorator

import "fmt"

type Food interface {
	Price() float64
	Description() string
}

// NOTE: x-burger decorator
type XBurger struct {
	Food Food
}

func (xb *XBurger) Price() float64 {
	price := 13.99

	if xb.Food != nil {
		return xb.Food.Price() + price
	}

	return price
}

func (xb *XBurger) Description() string {
	desc := "x-burger delicious"

	if xb.Food != nil {
		return xb.Food.Description() + fmt.Sprintf(", with %v", desc)
	}

	return desc
}

// NOTE: coke decorator
type Coke struct {
	Food Food
}

func (c *Coke) Price() float64 {
	price := 3.99

	if c.Food != nil {
		return c.Food.Price() + price
	}

	return price
}

func (c *Coke) Description() string {
	desc := "coke"

	if c.Food != nil {
		return c.Food.Description() + fmt.Sprintf(", with %v", desc)
	}

	return desc
}

// NOTE: sweet decorator
type Sweet struct {
	Food Food
}

func (s *Sweet) Price() float64 {
	price := 5.15

	if s.Food != nil {
		return s.Food.Price() + price
	}

	return price
}

func (s *Sweet) Description() string {
	desc := "sweet"

	if s.Food != nil {
		return s.Food.Description() + fmt.Sprintf(", with %v", desc)
	}

	return desc
}
