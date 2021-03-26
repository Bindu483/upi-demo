package main

import (
	"bufio"
	"fmt"
	"github.com/Bindu483/upi-demo/hdfc"
	"github.com/Bindu483/upi-demo/sbi"
	"github.com/Bindu483/upi-demo/upi"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Google Pay!!!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Please enter your mobile number: ")
	number, _ := reader.ReadString('\n')

	fmt.Print("Please chose your bank 1 for HDFC and 2 for SBI: ")
	bank, _ := reader.ReadString('\n')

	customer := &Customer{
		Name:         name,
		MobileNumber: number,
		Bank:         bank,
	}

	var upiBank upi.UPI

	//var stringVar string
	//
	//stringVar="sjbdfe"
	//stringVar=9876

	switch strings.TrimSpace(bank) {
	case "1":
		{
			//Liskov's substitution principle
			fmt.Println("Hey " + customer.Name + " you chose HDFC")
			upiBank = &hdfc.Account{
				BranchName:    "Mysore",
				AccountNumber: "987654345768767564",
				MobileNumber:  number,
				Balance:       0,
				Transactions:  []*upi.Transaction{},
			}
		}

	case "2":

		{
			fmt.Println("Hey " + customer.Name + " you chose SBI")

			upiBank = &sbi.Account{
				BranchName:    "Mysore",
				AccountNumber: "987654345768767564",
				MobileNumber:  number,
				Balance:       0,
				Transactions:  []*upi.Transaction{},
				FirstName:     name,
			}
		}

	}

	for {
		fmt.Println("************************************************************")
		fmt.Print("Please chose your transaction: 1 for credit and 2 for debit and 3 for balance: ")
		transactionType, _ := reader.ReadString('\n')

		var recipientNumber string
		var amountFloat float64

		if strings.TrimSpace(transactionType) != "3" {
			var err error

			fmt.Print("Please enter your recipient number: ")
			recipientNumber, _ = reader.ReadString('\n')
			fmt.Print("Please enter amount: ")
			amount, _ := reader.ReadString('\n')
			amount = strings.TrimSpace(amount)

			amountFloat, err = strconv.ParseFloat(amount, 64)
			if err != nil {
				fmt.Println("there is an error occured: ", err.Error())
				continue
			}
		}

		switch strings.TrimSpace(transactionType) {
		case "1":
			{
				_, err := upiBank.Credit(&upi.Transaction{
					SenderMobileNumber:   customer.MobileNumber,
					ReceiverMobileNumber: recipientNumber,
					Amount:               amountFloat,
					Type:                 upi.TransactionTypeCredit,
				})

				if err != nil {
					fmt.Println("error occured when doing credit")
					fmt.Println("error is :", err.Error())
					continue
				}
			}
		case "2":
			{
				_, err := upiBank.Debit(&upi.Transaction{
					SenderMobileNumber:   customer.MobileNumber,
					ReceiverMobileNumber: recipientNumber,
					Amount:               amountFloat,
					Type:                 upi.TransactionTypeDebit,
				})

				if err != nil {
					fmt.Println("error occured when doing debit")
					fmt.Println("error is :", err.Error())
					continue
				}
			}
		case "3":
			balance, err := upiBank.GetBalance()
			if err != nil {
				fmt.Println("error occured when getting balance")
				fmt.Println("error is :", err.Error())
				continue
			}

			fmt.Println("Your balance is :", balance.AccountBalance)
		}

	}
}

type Customer struct {
	Name         string
	MobileNumber string
	Bank         string
}

