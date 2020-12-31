package main

import "fmt"

func main(){//main method starts here
	fmt.Println("WElcome to Loop control statements")
//Long method
	fmt.Println("**** LONG METHOD ****")
	i := 1
	for i <=10{
		fmt.Printf("Long %d\t",i)
		i++
	}
	fmt.Printf("\n\n")

	//short method
	fmt.Println(" *** Short method ***")
	for i := 1;i <= 10;i++ {
		fmt.Printf("short %d\t",i)
	}
	fmt.Printf("\n\n")

	//FizzBuzz program for numbers between 1 -50
	/*
	If divisible by 3 "Fizz"
	if divisble by 5 " Buzz"
	if divisible by 3 and 5 then "FizzBuzz
	"*/
	for i:=1;i<=50;i++{
		if i % 15 == 0{
			fmt.Println("FizzBuzz")
		}else if i % 3== 0{
			fmt.Println("Fizz")
		}else if i % 5== 0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
			  
	}


}//main method end here