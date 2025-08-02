package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("hello world in go language")
	myutil.PrintMessage("satyendra call my utill package")
	var age int = 25
	var height float64 = 5.9
	var name string = "satyendra"
	var isAdult bool = true
	var version = 1.0
	const pi = 3.14159
	person := "person name"
	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("height", height)
	fmt.Println("is Adult", isAdult)
	fmt.Println("version", version)
	fmt.Println("Pi", pi)
	fmt.Println("person", person)
}
