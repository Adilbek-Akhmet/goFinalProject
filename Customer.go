package main

import "fmt"

type Customer struct {
	city, street, email, phone string
}

type CustomerBuilder struct {
	customer *Customer
}

func (c *CustomerBuilder) CustomerAddressInfo() *CustomerAddressBuilder {
	return &CustomerAddressBuilder{*c}
}

func (c *CustomerBuilder) CustomerContactInfo()  *CustomerContactBuilder {
	return &CustomerContactBuilder{*c}
}

func (c *CustomerBuilder) Build() *Customer {
	return c.customer
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{&Customer{}}
}

type CustomerAddressBuilder struct {
	CustomerBuilder
}

func (a *CustomerAddressBuilder) City (city string) *CustomerAddressBuilder {
	a.customer.city = city
	return a
}

func (a *CustomerAddressBuilder) Street (street string) *CustomerAddressBuilder {
	a.customer.street = street
	return a
}



type CustomerContactBuilder struct {
	CustomerBuilder
}

func (b *CustomerContactBuilder) Email(email string) *CustomerContactBuilder {
	b.customer.email = email
	return b
}

func (b *CustomerContactBuilder) Phone(phone string) *CustomerContactBuilder {
	b.customer.phone = phone
	return b
}



type CustomerPayment interface {
	pay()
}

type cash struct { }
func (c *cash) pay() {
	fmt.Println("Paid by cash!")
}

type cart struct{}
func (c *cart) pay() {
	fmt.Println("Paid by cart!")
}

type cartAdapter struct { cart *cart }

func (c *cartAdapter) pay() {
	fmt.Println("Adapting payment to cart!")
	c.cart.pay()
}

func(c *Customer) makePayment(money CustomerPayment) {
	money.pay()
}

