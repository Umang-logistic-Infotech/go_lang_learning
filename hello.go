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
	a, b := split(20)
	fmt.Printf("20 : %d %d\n", a, b)

	x, y := split(45)
	fmt.Printf("45 : %d %d\n", x, y)

	x, y = split(60)
	fmt.Printf("60 : %d %d\n", x, y)

	x, y = split(18)
	fmt.Printf("18 : %d %d\n", x, y)

}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}
