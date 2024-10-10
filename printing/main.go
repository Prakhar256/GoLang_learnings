package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Invalid")
	}
	return a / b, nil
}
func swapp(num1, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}

type Animal struct {
	LegCnt int
	color  string
}

func Hello() {
	fmt.Println("Hello ji! ki haal chaal")
	time.Sleep(9000. * time.Millisecond)
	fmt.Println("Kw hua gend phat gyi")
}
func Hi() {
	
	fmt.Println("Hi")
}
func main() {
	// var name string
	// fmt.Println("Hello what's your name")
	// fmt.Scan(&name)
	// fmt.Println("My name is", name)

	// scanner := bufio.NewReader(os.Stdin)
	// name, _ := scanner.ReadString('\n')
	// fmt.Println("Hello", name)

	// ans, _ := divide(10, 0)
	// fmt.Println(ans)

	// var arr [5]int
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// fmt.Println(len(arr))

	// Dynamic arrays -> slice just like vectors
	// vec := []int{1, 2, 3, 4, 5, 6}
	// vec = append(vec, 7, 8, 9)
	// fmt.Println(len(vec))

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }

	// i:=0;
	// for i<5 {
	// 	i++;
	// 	fmt.Println(i)
	// }

	// for i, v:=range vec{
	// 	fmt.Println(i,v);
	// }

	// arr1:=[]int {};
	// arr1=append(arr1, 4)

	// var n int
	// fmt.Scanln(&n)
	// var val int

	// vec1 := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	// fmt.Scanln(&val)
	// 	// vec1 = append(vec1, val)
	// 	fmt.Scanf("%d", &vec1[i])
	// }

	// for i := 0; i < n; i++ {
	// 	fmt.Print(vec1[i], " ")
	// }

	// vector := make([]string, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%s", &vector[i])
	// }

	// for _, value := range vector {
	// 	fmt.Print(value, " ")
	// }

	// mp := make(map[string]int)
	// for i:=0;i<n;i++{
	// 	mp[vector[i]]=i
	// }
	// for key,value:=range mp{
	// 	fmt.Println(key,value)
	// }
	// mp["Prakhar"]=1;
	// fmt.Println(mp["Prakhar"])

	// _, exists:= mp["Prakhar"]
	// fmt.Println(exists)

	var num1 int = 5
	var num2 int = 10
	swapp(&num1, &num2)
	// fmt.Println(num1, num2)

	var cow Animal
	cow.LegCnt = 4
	cow.color = "White"
	// fmt.Println(cow)

	// a1 := Animal{
	// 	LegCnt: 4,
	// 	color:  "Black",
	// }
	// fmt.Println(a1)
	buff := new(Animal)
	buff.LegCnt = 4
	buff.color = "Black"
	// fmt.Println(buff)

	// defer fmt.Println("Fuck you")
	// Defer keyword follows LIFO principle

	// That statement would execute just before returm 0 of main function
	// fmt.Println("Samay sab bataega")

	// File Handling
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Create or open output.txt for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Use buffered reading for input file
	scanner := bufio.NewScanner(inputFile)

	// Reading each line from input file and writing it to output file
	for scanner.Scan() {
		line := scanner.Text() // Get the text line by line from input file

		// Writing to output file
		_, err := fmt.Fprintln(outputFile, line) // Writing the line to output.txt
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// arr := make([]int, 5)
	// for i := 0; i < 5; i++ {
	// 	fmt.Scanf("d", &arr[i])
	// }
	// fmt.Println(arr)

	Hello()
	go Hi()
	// time.Sleep(800. * time.Millisecond)

}
