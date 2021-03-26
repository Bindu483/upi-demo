package upi
//this is defined by RBI
type UPI interface {
	Debit(tr *Transaction) (bool, error)
	Credit(tr *Transaction) (bool, error)
	GetBalance() (*Balance, error)
}

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "Debit"
	TransactionTypeCredit TransactionType = "Credit"
)

type Transaction struct {
	SenderMobileNumber   string
	ReceiverMobileNumber string
	Amount               float64
	Type                 TransactionType
}

type Balance struct {
	MobileNumber   string
	AccountBalance float64
}
