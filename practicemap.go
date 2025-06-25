package main

import ("fmt"
)

func main3() {
	menu := map[string]float64{
		"cake":    7.90,
		"bun":     8.90,
		"soup":    9.80,
		"parrota": 9.70,
	}
	fmt.Println(menu)
	fmt.Println(menu["cake"])
	var n float64
	// using loop in the map
	for k, v := range menu {
		fmt.Println(k, "-", v)
		n += v
		fmt.Println("total value of the menu", n)

	}

	// map using ints

	phone := map[int]string{
		1: "praveen",
		2: "rajesh",
		3: "saswant",
		4: "ramesh",
	}
	fmt.Println(phone)
	fmt.Println(phone[3])
	for i, j := range phone {
		fmt.Println("roll no:", i, " ", "name:", j)
	}
}
