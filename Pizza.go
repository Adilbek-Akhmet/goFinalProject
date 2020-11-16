package main

type pizza interface {
	getPrice() int
}

type habanero struct{}

func (h *habanero) getPrice() int {
	return 1500
}

type hawaii struct{}

func (h *hawaii) getPrice() int {
	return 2500
}

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 700
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 500
}


