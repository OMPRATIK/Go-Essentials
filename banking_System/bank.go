package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read the balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse tbe balance")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println("Welcome to Go Bank !")

	/*
		// infinite for loop
		for {

		}
	*/
	// ---------------------------------
	/*
		// for loop
		for i := 1; i < 10; i++ {
			fmt.Println("Hello")
		}
	*/
	for {
		fmt.Println("what do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositValue float64
			fmt.Scan(&depositValue)
			if depositValue <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			accountBalance += depositValue
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawlValue float64
			fmt.Scan(&withdrawlValue)
			if withdrawlValue > accountBalance {
				fmt.Println("Insufficient Balance : ")
				fmt.Println("Your balance :", accountBalance)
			} else if withdrawlValue <= 0 {
				fmt.Println("Invalid amount!")
				continue
			} else {
				fmt.Printf("%f is withdrawn", withdrawlValue)
				accountBalance -= withdrawlValue
				writeBalanceToFile(accountBalance)
				fmt.Println("Your balance updated! :", accountBalance)
			}
		} else if choice == 4 {
			fmt.Println("Thank You for yout visit!")
			break
		} else {
			fmt.Println("Not a valid option!")
		}
	}
	/*
		// switch statement (no break needed)
		choice := 10
		switch choice {
		case 1:
			fmt.Println(choice)
		case 2:
			fmt.Println(choice);
		default:
			fmt.Println("Good Bye !")
		}
	*/
}
