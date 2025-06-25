package main

type bill struct {
	name  string
	iteam map[string]float64
	tip   float64
}

// new bill for do
func newbill(name string) bill {
	b := bill{
		name:  name,
		iteam: map[string]float64{"cake":    7.90,
		"bun":     8.90},
		tip:   0,
	}
	return b
}


// function to relate 
func(b bill) format(){

}
