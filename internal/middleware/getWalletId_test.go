package middleware

import (
	"testing"

)

func TestGetWalletInfo(t *testing.T) {

	result := queryWallet(20001)

	if result != "20001|7 - 14|1234" {
		t.Error("result != 20001|7 - 14|1234 ",result)
	}

}

