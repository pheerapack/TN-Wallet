package middleware

import (
	"strconv"
)

/*
type W interface {
	Balance() int
}


type realWallet_Id struct{

}

type fakeWallet_Id struct{
	balance int
}

func (r realWallet_Id) Balance(i int) int {
	return getDbBal()
}


func (f fakeWallet_Id) Balance(i int) int {
	return f.balance
}

type WalletAccount struct {
	W
}
*/
/*

func DebitAccount (currentBalance int,TrnAmt int) string{
	var endBal int
	var txnResult string
	var statusCode int
	var errMsg string

	errMsg=""
	statusCode=0

	endBal=currentBalance-TrnAmt
	txnResult=strconv.Itoa(statusCode)+"|"+strconv.Itoa(endBal)+"|"+errMsg
	// txn_status 0 success 1 fail|endBal|error desc (if error)
	return txnResult
}

*/
type  WalletAccount struct {
	walletId string
	balance int
}

func  (w *WalletAccount)DebitAccount (TrnAmt int) string{
	var txnResult string
	var statusCode int
	var errMsg string

	errMsg=""
	if w.balance>=TrnAmt {
		w.balance=w.balance-TrnAmt
	} else {
		errMsg="Not enough money"
		statusCode=1
	}

	txnResult=strconv.Itoa(statusCode)+"|"+errMsg
	// txn_status 0 success 1 fail|endBal|error desc (if error)
	return txnResult
}
