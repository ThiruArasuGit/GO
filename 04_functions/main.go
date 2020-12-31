package main

import "fmt"

func greeting(name string) string{
	return "Hello " +name
}

func getsum(a int , b int) int {
	return (a+b)
}

func getsumoffloats (num1 , num2 float32) float32{
	return num1 + num2
}


//************************This is the main function to be executed **********************
func main(){

	fmt.Println("************This is Functions******************")
	fmt.Println(greeting("Thiru"))
	fmt.Println(getsum(10,20))
	fmt.Println(getsumoffloats(1.101,2.3))

}