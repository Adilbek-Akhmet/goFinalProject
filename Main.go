package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main() {
	fmt.Println("---Pizza Delivery Service---")
	fmt.Println("1-Order")
	fmt.Println("2-Company Info")
	fmt.Println("3-For Workers")
	fmt.Println("----------------------------")
	var p int
	fmt.Scanln(&p)
	switch {
	case p == 1:
		fmt.Println("Pizza")

		fmt.Println("1-Hawaii")
		fmt.Println("2-Hawaii with Cheese Topping")
		fmt.Println("3-Hawaii with Tomoto Topping")
		fmt.Println("4-Hawaii with Cheese and Tomoto Topping")
		fmt.Println("-----------")
		fmt.Println("5-Habanero")
		fmt.Println("6-Habanero with Cheese Topping")
		fmt.Println("7-Habanero with Tomoto Topping")
		fmt.Println("8-Habanero with Cheese and Tomoto Topping")

		fmt.Scanln(&p)
		switch {
		case p == 1:
			pizza := &hawaii{}
			fmt.Printf("It costs %d\n", pizza.getPrice())
		case p == 2:
			pizza := &hawaii{}
			pizzaWithCheese := &cheeseTopping{pizza}
			fmt.Printf("It costs %d\n", pizzaWithCheese.getPrice())
		case p == 3:
			pizza := &hawaii{}
			pizzaWithTomoto := &cheeseTopping{pizza}
			fmt.Printf("It costs %d\n", pizzaWithTomoto.getPrice())
		case p == 4:
			pizza := &hawaii{}
			pizzaWithCheese := &cheeseTopping{pizza}
			pizzaWithCheeseAndTomoto := &tomatoTopping{pizzaWithCheese}
			fmt.Printf("It costs %d\n", pizzaWithCheeseAndTomoto.getPrice())
		case p == 5:
			pizza := &habanero{}
			fmt.Printf("It costs %d\n", pizza.getPrice())
		case p == 6:
			pizza := &habanero{}
			pizzaWithCheese := &cheeseTopping{pizza}
			fmt.Printf("It costs %d\n", pizzaWithCheese.getPrice())
		case p == 7:
			pizza := &habanero{}
			pizzaWithTomoto := &cheeseTopping{pizza}
			fmt.Printf("It costs %d\n", pizzaWithTomoto.getPrice())
		case p == 8:
			pizza := &habanero{}
			pizzaWithCheese := &cheeseTopping{pizza}
			pizzaWithCheeseAndTomoto := &tomatoTopping{pizzaWithCheese}
			fmt.Printf("It costs %d\n", pizzaWithCheeseAndTomoto.getPrice())
		}
		Customer := &Customer{}
		fmt.Println("Do you want to pay by 1-cash or 2-cart")
		fmt.Scanln(&p)
		switch {
		case p == 1:
			cash := &cash{}
			Customer.makePayment(cash)
		case p == 2:
			cart := &cart{}
			cartAdapter := &cartAdapter{cart: cart}
			Customer.makePayment(cartAdapter)
		}
		fmt.Println("-----------")
		var city string
		var street string
		var email string
		var phone string
		fmt.Println("Customer Address Info")
		fmt.Println("Enter your city")
		fmt.Scanln(&city)
		fmt.Println("Enter your street")
		fmt.Scanln(&street)
		fmt.Println("Customer Contact Info")
		fmt.Println("Enter your email")
		fmt.Scanln(&email)
		fmt.Println("Enter your phone number")
		fmt.Scanln(&phone)
		customer := NewCustomerBuilder()
		customer.CustomerAddressInfo().City(city).Street(street)
		customer.CustomerContactInfo().Email(email).Phone(phone)
		customerInfo := *customer.Build()
		fmt.Println("You will get your order at this address: " + customerInfo.city + ", " + customerInfo.street)
		fmt.Printf("We contact you by these info: email %v, phone-number %v\n", customerInfo.email, customerInfo.phone)
		p := strconv.Itoa(p)
		lines := [4]string{customerInfo.email, "Pizza index: " + p,"Address Info: " + customerInfo.city + ", " + customerInfo.street,"Phone-number: " + customerInfo.phone}
		f, err := os.OpenFile("order.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
			w := bufio.NewWriter(f)
		for _, s := range lines {
			w.WriteString(s + "\n")
		}
		w.Flush()
	case p == 2:
		fmt.Println("Founder of our company")
		directorFactory := NewEmployeeFactory("director")
		director1 := directorFactory("Bagdaulet Alimkhan")
		director2 := directorFactory("Adilbek Akhmet")
		fmt.Println("Name: " + director1.Name + " Position: " + director1.Position)
		fmt.Println("Name: " + director2.Name + " Position: " + director2.Position)
		fmt.Println("Contact Info")
		fmt.Println("Phone: +7885545")
		fmt.Println("Email: pizza@gmail.com")
	case p == 3:
		db := GetSingletonDatabase()
		db.showOrderList()
		fmt.Println("1-Search")
		fmt.Println("2-Exit")
		fmt.Scanln(&p)
		switch {
		case p == 1:
			fmt.Println("Enter email of customer")
			var email string
			fmt.Scanln(&email)
			fmt.Println(db.search(email))
		case p == 2:
			os.Exit(0)
		}
	}
}
