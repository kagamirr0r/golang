package main

import "fmt"

type Payment interface {
	Pay(amount int)
}

// Cash
type Cash struct{}

func (c Cash) Pay(amount int) {
	fmt.Printf("%d円を現金で支払う\n", amount)
}

// Suica
type Suica struct {
	Deposit int
}

func (s Suica) Pay(amount int) {
	fmt.Printf("%d円をSuicaで支払う\n", amount)
}

func (s *Suica) Charge(amount int) {
	s.Deposit += amount
	fmt.Printf("%d円をチャージする\n", amount)
}

// main
func main() {
	var cash Payment = Cash{}
	var suica Suica = Suica{}

	pa := []Payment{cash, suica}
	for _, payment := range pa {
		payment.Pay(5000)
	}

	suica.Charge(2000)

	suica2 := &Suica{Deposit: 1500}
	suica2.Charge(1500)
	suica2.Charge(1500)
	fmt.Println(suica2.Deposit)
}
