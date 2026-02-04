package main

import "fmt"

func main() {

	var name string = "Umang"
	name2 := "Umang"
	fmt.Println("Hello, Go!", name, name2)

	var age int = 25
	var percentage float32 = 245.45181561
	fmt.Println(age, percentage)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", len(arr))
}
