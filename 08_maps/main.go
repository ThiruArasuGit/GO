package main

import "fmt"

func main(){
	fmt.Println(" *** Welcome to Maps **")

//Map is key and value pair data type
//Method -1
	//declare map variable
	emails_1 := make(map[string]string)

	//Assign key - value to map variable
	emails_1["T"] = "thiru@gmail.com"
	emails_1["A"] = "anitha@gmail.com"
	emails_1["G"] = "gsg@gmail.com"
	emails_1["S"] = "shweta@gmail.com"

	fmt.Println(emails_1)     //Prints all the values with K-V
	fmt.Println(emails_1["T"])//Prints specific K-V
	fmt.Println(len(emails_1))//Prints size of the map values

	//delete a value from map variable
	delete(emails_1, "T") //Note down the syntax of deletion

	fmt.Println("Final list after deleting the Key = 'T': ")
	fmt.Println(emails_1)

//Method -2 
    //Declare map variable
   emails_2 := map[string]string{"MS":"https://metricstream.com","JDA":"https://jda.com","SG":"https://socgen.com","DW":"https://dataweaves.com"}

   fmt.Println(emails_2)
   fmt.Println(len(emails_2))


}