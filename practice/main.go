package main

import (
	"fmt"
)

const aConstant int = 100

func main() {

	var aString string = "This is Go Lang"
	fmt.Println(aString)
	fmt.Printf("the varible's type is %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Println("the variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("the variable's type is %T\n", anotherInt)

	mystring := "This is a string with short declaration"
	fmt.Println(mystring)
	fmt.Printf("the variable's type is %T\n", mystring)

	fmt.Println(aConstant)
	fmt.Printf("the variable's type is %T\n", aConstant)

	test

}
