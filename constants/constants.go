package main

import "fmt"

func sum(a int ,b int)int{
	return a+b
}
func main(){
	const a = 10
const b = 20
c := sum(a,b)
	fmt.Println("the value of a is",a)
	fmt.Println("the value of b is",b)
	fmt.Println("the value of c is",c)

}