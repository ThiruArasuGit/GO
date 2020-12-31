package main

import "fmt"
//Closures are anonymous function

//Adder function returns anonymous function ************** DO some more on this *****************
func adder() func(int) int{
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main(){

	fmt.Println("-- welcome to Closures ---")

	sum := adder()

	for i:=0;i<=10;i++{
		fmt.Println(sum(i))
	}
}