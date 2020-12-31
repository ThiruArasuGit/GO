package main

import "fmt"


//********* This is main function

func main(){
	fmt.Println("Welcome to 05_ARRAYS_AND_SLICES")

	//Declare and asisgn values to array
	//Method -1
	var fruitArr1[2]string
	fruitArr1[0] = "Apple"
	fruitArr1[1] = "Orange"
	fmt.Println( fruitArr1)

	//Method 2
	fruitArr2 :=[2]string{"Banana", "Grapes"}
	fmt.Println(fruitArr2)
// Note: ARRAY must be delcared with fixed SIZE.

fruitSlice :=[]string{"Mango","PineApple","Supporta"}
fmt.Println(fruitSlice)
fmt.Println(len(fruitSlice))
fmt.Println(fruitSlice[0:2])
//Note: Slice does not have fixed size

}