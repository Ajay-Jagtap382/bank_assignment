package main

import (
	acc "bank/admin"
	cus "bank/user"
	"fmt"
)

func main() {
	fmt.Println("Hello , Enter the number which suits your role ")

	for {
		fmt.Println("1 .Accountant ")
		fmt.Println("2 .Customer ")
		fmt.Println("3 .Exit ")

		fmt.Println("Enter the number : ")
		choice := -1
		fmt.Scanln(&choice)

		if choice == 1 {
			acc.Accountant()
		} else if choice == 2 {
			cus.Custo()
		} else if choice == 3 {
			break
		} else {
			fmt.Println("Incorrect Input!! Try again!!")
		}
	}
}
