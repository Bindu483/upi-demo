package hdfc

import (
	"errors"
	"github.com/Bindu483/upi-demo/upi"
)

type Account struct {
	BranchName    string
	AccountNumber string
	MobileNumber  string
	Balance       float64
	Transactions  []*upi.Transaction
}

func (a *Account) Debit(tr *upi.Transaction) (bool, error) {
	if a.Balance-tr.Amount < 0 {
		return false, errors.New("not sufficient funds")
	}

	a.Balance = a.Balance - tr.Amount
	a.Transactions = append(a.Transactions, &upi.Transaction{
		SenderMobileNumber:   tr.SenderMobileNumber,
		ReceiverMobileNumber: tr.SenderMobileNumber,
		Amount:               tr.Amount,
		Type:                 upi.TransactionTypeDebit,
	})

	return true, nil
}

func (a *Account) Credit(tr *upi.Transaction) (bool, error) {
	a.Balance += tr.Amount

	a.Transactions = append(a.Transactions, &upi.Transaction{
		SenderMobileNumber:   tr.SenderMobileNumber,
		ReceiverMobileNumber: tr.SenderMobileNumber,
		Amount:               tr.Amount,
		Type:                 upi.TransactionTypeCredit,
	})

	return true, nil
}

func (a *Account) GetBalance() (*upi.Balance, error) {
	return &upi.Balance{
		MobileNumber:   a.MobileNumber,
		AccountBalance: a.Balance,
	}, nil
}

