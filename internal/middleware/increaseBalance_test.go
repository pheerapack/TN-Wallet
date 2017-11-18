package middleware

import "testing"

func TestDisplayMerchantBalanceAfterIncrease1000(t *testing.T) {
	m := merchant{
		Wallid:  "20001",
		Balance: 1000,
	}
	m.Increase(1000)
	if m.Balance != 2000 {
		t.Error("Balance != 2000")
	}

}


