package main

import "fmt"

func main(){
	i :=1
	for i<10{
		fmt.Println("the value of i is",i)
		i++
	}

	for i:= range 3{
		fmt.Println("the value of i is",i)
	}
}