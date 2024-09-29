package main

// idhar class nhi hoti package hota hai

// files are tightly coupled in go language so jo bhi files ek folder me padi hai wo sab ek package me treat ki jaati hai like for example if there is another file in GOLANG folder and we assigned it package other than main then it will show error so to assign other package then we have to make another folder inside GOLANG folder and keep that file into it.

// Ek directory me ek hi tarah ka package use kar sakte hai , agar diffrent package use kane hai sub-directory bana lo

import (
	"fmt"
)

func main() {
	// main function is a part of package main, koi dusra package hota like hello forld to usme main function nhi hota

	// fmt.Println("Hello World ")
	// fmt.Print("Kya hua ")
	// myLearnings.PrintMsg("Prakhar")
	// kuch bhi auto import ke liye proxy or ethernet off karo

	var str string = "Starc"
	fmt.Println(str)
	temp := 21
	fmt.Println(temp)
	fmt.Printf("My name is %s", str)

}
