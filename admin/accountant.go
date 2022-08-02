package Acc

import (
	"fmt"
	"math/rand"
	"strings"
)

type customer struct {
	email       string
	password    int
	phone_num   string
	account_num int
	balance     int
}

var data = make(map[int]customer)

var i = 1

func create(key customer) {
	email := ""
	fmt.Println("Enter customer Email Id : ")
	fmt.Scanln(&email)
	key.email = email
	phone_n := ""
	fmt.Println("Enter customer Phone number : ")
	fmt.Scanln(&phone_n)
	key.phone_num = phone_n

	key.balance = 0
	key.account_num = i
	i++

	key.password = rand.Intn(100000000) + i

	data[key.account_num] = key

	fmt.Println("Account Number : ", key.account_num)
	fmt.Println("Account Password : ", key.password)
	fmt.Println("Account created")
}

func del(acc int) {
	_, ok1 := data[acc]
	if ok1 {
		delete(data, acc)
		fmt.Println("Account Deleted Successfully ")
	} else {
		fmt.Println("Account does not exist ,try again! ")
	}
}

func verify() {
	Acc_email := "account@bank.com"
	Acc_password := "josh@123"

	input_email := ""
	input_password := ""

	fmt.Print("Email : ")
	fmt.Scanln(&input_email)

	input_email = strings.ToLower(input_email)

	if input_email != Acc_email {
		fmt.Println("Wrong Email address , Try again! ")
		return
	}
	fmt.Print("Password : ")
	fmt.Scanln(&input_password)

	if input_password != Acc_password {
		fmt.Println("Wrong password , Try again! ")
		return
	}

}

func Accountant() {

	verify()

	for {
		n := -1
		fmt.Println(" 1. Create the account ")
		fmt.Println(" 2. Delete the account ")
		fmt.Println(" 3. Exit ")
		fmt.Print(" Choose the number : ")
		fmt.Scanln(&n)

		if n == 1 {
			var cus1 customer
			create(cus1)
		} else if n == 2 {
			acc := 0
			fmt.Println(" Enter the account number : ")
			fmt.Scanln(&acc)
			del(acc)
		} else if n == 3 {
			break
		} else {
			fmt.Println("Invalid Input ")
			continue
		}
	}
}

// func Expo() map[int]customer {
// 	return data
// }

func Getpassword(acco int) int {
	return data[acco].password
}

func Getbalance(acco int) int {
	return data[acco].balance
}

func Verify_acc(acco int) bool {
	_, ok1 := data[acco]
	if ok1 {
		return true
	} else {
		return false
	}
}

func Verify_pass(acco int, pass int) bool {
	var key = data[acco]
	if key.password == pass {
		return true
	} else {
		return false
	}

}

func Add_cash(acco int, cash int) {
	key := data[acco]
	key.balance += cash
	data[acco] = key
}

func Remove_cash(acco int, amount int) bool {
	key := data[acco]
	curr_balance := key.balance
	if curr_balance < amount {
		fmt.Println("Withdraw amount is greater than account balance, cannot withdraw the amount ")
		fmt.Println("Your current balance is : ", key.balance)
		fmt.Println("Please enter the amount less than balance")
		return false
	} else {
		key.balance -= amount
		data[acco] = key
		return true
	}
}

func Show_bal(acco int) int {
	return data[acco].balance
}
