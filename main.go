package main

import "fmt"

func start() {

	fmt.Println("Enter the number : ")
	choice := -1
	fmt.Scanln(&choice)

	if choice == 1 {

	} else if choice == 2 {

	} else {
		fmt.Println("Incorrect Input!! Try again!!")
		start()
	}
}

func main() {
	fmt.Println("Hello , Enter the number which suits your role ")
	fmt.Println("1 .Accountant ")
	fmt.Println("2 .Customer ")

}
