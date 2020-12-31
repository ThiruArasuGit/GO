package main

import "fmt"


func main(){//Main starts here
	fmt.Println("Welcome to conditionals statements")

	x:=15
	y:=10

	//If..else 
	fmt.Println("------- If ELSE statement ---------------")
	if x <y {
		fmt.Printf("%d is less than %d\n",x,y)
	}else{
		fmt.Printf("%d is less than %d\n",y,x)
	}

	//if..else if..
	fmt.Println("------- ELSE IF statement ---------------")
	color := "red1"

	if color == "red"{
		fmt.Println("Code is red")
	}else if color == "green"{
		fmt.Println("Code is green")
	}else{
		fmt.Println("Code is neither red not green")
	}


//switch statement
fmt.Println("------- Switch statement ---------------")
switch color{
case "red":
	fmt.Println("Code is red")
case "green":
	fmt.Println("Code is green")
default:
	fmt.Println("Code is neither red nor green")

}







}//Main ends here