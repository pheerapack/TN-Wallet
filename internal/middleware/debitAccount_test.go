package middleware

import (
	"testing"
	"fmt"
	"strconv"
)

func TestDebitSuccess(t *testing.T) {

	w:=WalletAccount{
		walletId:  "10002",
		balance: 2000,
	}

	txnResult:=w.DebitAccount(1000)
	if ((txnResult!="0|")||(w.balance!=1000)) {
		fmt.Println("result is not true, value is "+txnResult)
	} else {
		// correct
		fmt.Println("Success case: Balance is "+ strconv.Itoa(w.balance))

	}
}

func TestDebitFailNotEnough(t *testing.T) {

	w:=WalletAccount{
		walletId:  "10002",
		balance: 2000,
	}

	txnResult:=w.DebitAccount(3000)
	if (txnResult!="1|Not enough money")||(w.balance!=2000) {
		fmt.Println("result is not true, value is "+txnResult)
	} else {
		// correct
		fmt.Println("Faile case : Balance is "+ strconv.Itoa(w.balance))
	}
}
