package main

import "fmt"

func main() {
	var (
		a=5
		b=10
		c=15
	)

	const x string = "Hello World"
	//x = "Some other string"
	fmt.Println(x)
	fmt.Println(a,b,c)
}