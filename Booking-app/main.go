package main

import ("fmt")

func main() {
	var num1 int
	fmt.Print("Enter a number:")
	fmt.Scan(&num1)

	switch num1{
	case 1:fmt.Println("Sunday")
	case 2:fmt.Println("Monday")
	case 3:fmt.Println("Tuesday")
	case 4:fmt.Println("Wednesday")
	case 5:fmt.Println("Thursday")
	case 6:fmt.Println("Friday")
	case 7:fmt.Println("Saturday")
	default:fmt.Println("Invalid input")
	}
}
