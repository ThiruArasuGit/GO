package main

import "fmt"

func main(){
	fmt.Println(" ** WElcome to RANGE **")

	//Range using slices
	nlist :=[]int{10,20,30,40}

	//Get the above nlist(slice) using index number
	fmt.Println("------- Print nlist with index --------")
	for idx, val := range nlist{
		fmt.Printf("%d - ID: %d\n", idx,val)
	}
	
	//Get the above nlist(slice) without index number
	fmt.Println("------- Print nlist without index --------")
	for _, val:=range nlist{
		fmt.Printf("ID - %d\n",val)
	}

	//Range using map
	employee :=map[int]string{101:"Thiru A",102:"Shweta M",103:"Anitha K",104:"Gsg G",105:"Roshan K"}

	//Print the above map values using range function
	fmt.Println("------- Print employee(map) --------")
	for k,v :=range employee{
		fmt.Printf("%d : %s\n",k,v)
	}

	fmt.Println("------- Print employee(map) keys --------")
	for k := range employee{
		//fmt.Println(k)
		fmt.Printf("Key: %d\n",  k)
	}
	fmt.Println("------- Print employee(map) names --------")
	for _,v:=range employee{
		fmt.Printf("Name: %s\n",v)
	}
}