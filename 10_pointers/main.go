package main

import "fmt"
//Pointers are the address of a value
func main(){

	fmt.Println(" *** Welcome to Pointers ***")

	a := 5
	b := &a   //b variable is assigned with the address of a which means a and b values are 5 (figuratively)

	//fmt.Println(a + b)// If we try to add link this it will thorugh an error as mismatched data type as b is address of a(not actual value of a)
	fmt.Println(a + *b) // Here we have to use * before b to invoke the actuall assigned value of b.


	//Check the type of a and b variavles to get the detailed explanation
	fmt.Printf("a type: %T\n",a)
	fmt.Printf("b type: %T\n",b) //returns as *int - Pointer data type

	fmt.Printf("a value: %d\n",a)
	fmt.Printf("b value: %d\n",b)








}