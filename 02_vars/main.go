package main
import "fmt"


func main(){
	var name string = "Thirunavukkarasu"
	var age = "30"
	var email = "thirunavukkarasu143@gmail.com"
	var phone = 9008756266
	
	fmt.Println(name + "\n"+ age + "\n"+ email+"\n")

	fmt.Println(phone)
	//Find the type of the variable
	fmt.Printf("%T\n",phone)

	var isCool = true
	fmt.Printf("%T\n",isCool)

	const isMe = false
	//isMe = true ** Can not overwrite const variable values as it is constant.
	fmt.Printf("%T\n", isMe)

	//Shorthand Declartion. We can not do this type of declartion outside a function.
	myName, myAge := "Thiru A", 31
	myMobile := 9008756266
	fmt.Println(myName,myAge,myMobile)

}