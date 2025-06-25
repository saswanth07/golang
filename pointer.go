package main

import ("fmt"
)

func updatename(x *string)  {
	*x = "rajesh"
}

func main2()  {
		s1 := "praveen"
     //updatename(s1)
	//fmt.Println("after calling the function :", s1)
	//fmt.Println(&s1)
	m:=&s1
	fmt.Println(m)

	// getting the memory address of the value 
	fmt.Println("value of the memory address :",*m)
	updatename(m)
	fmt.Println(s1)

}