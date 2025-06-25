package main

import ("fmt"
)

func updatename1(x string) string {
	x = "rajesh"
	return x
}
func updatemenu(y map[string]float64) {
	y["coffee"] = 9.99
}
func main1() {
	s1 := "praveen"
	fmt.Println("before calling function :", s1)
	s1 = updatename1(s1)
	fmt.Println("after calling the function :", s1)

	menu := map[string]float64{
		"cake":    7.90,
		"bun":     8.90,
		"soup":    9.80,
		"parrota": 9.70,
	}
	updatemenu(menu)
	fmt.Println(menu)
}
