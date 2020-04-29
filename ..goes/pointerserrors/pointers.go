package pointerserrors

import "fmt"
import "errors"

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(newAmount Bitcoin) Bitcoin {
	fmt.Printf("wallet deposit %v \n", w)
	w.balance += newAmount
	return w.balance
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("wallet balance  %v \n", w.balance)
	return w.balance
}

func (w *Wallet) WithDraw(newAmount Bitcoin) error {
	fmt.Printf("wallet withdrawal %v \n", w)
	if newAmount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= newAmount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
