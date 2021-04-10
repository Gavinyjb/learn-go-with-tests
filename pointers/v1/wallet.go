package pointers

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}
func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}
