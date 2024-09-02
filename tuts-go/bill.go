package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func (b *bill) formatBill() string {
	fs := "bill break-down : \n\n"
	var total float64
	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v %-10s £%v \n", k, "------", v)
		total += v
	}

	fs += fmt.Sprintf("\n%-10s £%0.2f\n", "tip :", b.tip)
	total += b.tip
	fs += fmt.Sprintf("\n%s £%0.2f\n", "total :", total)
	return fs
}
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}
func createNewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) saveBill() {
	data := []byte(b.formatBill())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved successfully")
}
