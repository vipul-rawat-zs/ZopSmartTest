package main

import (
	"fmt"
	"strconv"
)

type order struct {
	id              int
	items           []string
	shippingFee     float64
	includesAlcohol bool
	amount          float64
}

func New(id string, items []string, amount float64) *order {
	idInt, _ := strconv.Atoi(id)
	o := order{id: idInt, items: items, amount: amount}
	return &o
}

func (o *order) Recalculate() {
	if o.amount < 250.0 {
		o.shippingFee = 10
	}
}

func main() {
	o := New("1022", []string{"dairy milk", "bounty", "milk"}, 240)
	fmt.Printf("id : %v, items : %v, amount : %v, shipping fee : %v\n", o.id, o.items, o.amount, o.shippingFee)

	o.Recalculate()
	fmt.Printf("id : %v, items : %v, amount : %v, shipping fee : %v\n", o.id, o.items, o.amount, o.shippingFee)
}
