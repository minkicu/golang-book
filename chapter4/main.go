package main

import "fmt"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "\nHello World"
	fmt.Println(y)

	var z string
	z = "\nfirst"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)

	var s string
	s = "\nfirst"
	fmt.Println(z)
	s = s+"second"
	fmt.Println(s)

	fmt.Println("\n")
	var s1 string = "hello"
	var s2 string = "world"
	fmt.Println(s1==s2)

	var s3 string = "hello"
	var s4 string = "hello"
	fmt.Println(s3==s4)

	s5 := "\nHello World"
	fmt.Println(s5)

	dogsName := "Max"
	fmt.Println("\nMy dog's name is", dogsName)

}