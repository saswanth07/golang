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

// new bill making
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
func (b *bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64
	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v...$%v\n", k+":", v)
		total += v
	}
	//tip
	fs += fmt.Sprintf("%-25v...$%v\n", "tip:", b.tip)
	//total
	fs += fmt.Sprintf("%-25v...$%0.2f", "total:", total+b.tip)
	return fs
}
func (b *bill) updatetip(tip float64) {
	b.tip = tip
}
func (b *bill) itemsadd(name string, price float64) {
	b.items[name] = price
}
func (b *bill) save(){
	data:=[]byte(b.format())
	err:=os.WriteFile("bills/"+b.name+".txt",data,0644)
	if(err!=nil){
		panic(err)
	}
	fmt.Println("bill saved to file")
}
