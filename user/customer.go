package Cust

import (
	acc "bank/admin"
	"fmt"
	"time"
)

type history struct {
	Sr_no              int
	Date               string
	Transaction_Type   string
	Transaction_amount int
	Total_balance      int
}

var his = map[int][]history{}

var dt = time.Now()

func Deposit(account_num int) {
	fmt.Println("Enter the amount you want to deposit : ")
	amount := 0
	fmt.Scanln(&amount)
	acc.Add_cash(account_num, amount)
	_, ok1 := his[account_num]
	if ok1 {
		var slice1 []history
		slice1 = his[account_num]
		var last_trans history = slice1[len(slice1)-1]
		last_sr := last_trans.Sr_no

		var cus_his1 history
		cus_his1.Sr_no = last_sr + 1
		cus_his1.Date = dt.Format("01-02-2006")
		cus_his1.Transaction_Type = "Credit"
		cus_his1.Transaction_amount = amount
		cus_his1.Total_balance = acc.Getbalance(account_num)
		slice1 = append(slice1, cus_his1)
		his[account_num] = slice1
	} else {
		var cus_his history
		cus_his.Sr_no = 1
		cus_his.Date = dt.Format("01-02-2006")
		cus_his.Transaction_Type = "Credit"
		cus_his.Transaction_amount = amount
		cus_his.Total_balance = acc.Getbalance(account_num)
		var slice []history
		slice = append(slice, cus_his)
		his[account_num] = slice
	}
}

func Withdraw(account_num int) {
	fmt.Println("Enter the amount you want to Withdraw : ")
	amount := 0
	fmt.Scanln(&amount)
	if acc.Remove_cash(account_num, amount) {

		_, ok1 := his[account_num]
		if ok1 {
			var slice1 []history
			slice1 = his[account_num]
			var last_trans history = slice1[len(slice1)-1]
			last_sr := last_trans.Sr_no

			var cus_his1 history
			cus_his1.Sr_no = last_sr + 1
			cus_his1.Date = dt.Format("01-02-2006")
			cus_his1.Transaction_Type = "Debit"
			cus_his1.Transaction_amount = amount
			cus_his1.Total_balance = acc.Getbalance(account_num)
			slice1 = append(slice1, cus_his1)
			his[account_num] = slice1
		} else {
			var cus_his history
			cus_his.Sr_no = 1
			cus_his.Date = dt.Format("01-02-2006")
			cus_his.Transaction_Type = "Debit"
			cus_his.Transaction_amount = amount
			cus_his.Total_balance = acc.Getbalance(account_num)
			var slice []history
			slice = append(slice, cus_his)
			his[account_num] = slice
		}
	}
}

func show_balance(account_num int) {
	bal := acc.Show_bal(account_num)

	fmt.Println("Your balance is : ", bal)
}

func show_history(account_num int) {
	var slice2 []history = his[account_num]

	for i := len(slice2) - 1; i >= 0; i-- {
		var temp history = slice2[i]
		fmt.Println("Sr_no is : ", temp.Sr_no)
		fmt.Println("Date is : ", temp.Date)
		fmt.Println("Transaction_Type is : ", temp.Transaction_Type)
		fmt.Println("Transaction_amount is : ", temp.Transaction_amount)
		fmt.Println("Total_balance is : ", temp.Total_balance)
	}
}

func Custo() {
	account_no, password := 0, 0
	fmt.Println("Enter the acount number ")
	fmt.Scanln(&account_no)

	if !acc.Verify_acc(account_no) {
		fmt.Println("Account number does not exists , try again!")
		return
	}

	fmt.Println("Enter the password ")
	fmt.Scanln(&password)

	choice := 0

	if acc.Verify_pass(account_no, password) {

		for {
			fmt.Println("1 .Deposit ")
			fmt.Println("2 .Withdraw ")
			fmt.Println("3 .See the balance ")
			fmt.Println("4 .see the Transaction History")
			fmt.Println("5 .Exit ")
			fmt.Scanln(&choice)

			if choice == 1 {
				Deposit(account_no)
			} else if choice == 2 {
				Withdraw(account_no)
			} else if choice == 3 {
				show_balance(account_no)
			} else if choice == 4 {
				show_history(account_no)
			} else if choice == 5 {
				break
			} else {
				fmt.Println("Incorrect Input!! Try again!!")
			}
		}
	} else {
		fmt.Println("wrong password , try again!")
		return
	}
}
