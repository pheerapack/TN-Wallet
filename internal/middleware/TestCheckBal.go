package middleware

import (
	"testing"
)

type wallet struct {
	WalletID int
	Bal float32
}


func TestGetCheckBalLess(t *testing.T) {
	result := CheckBal(10002,1000)

	if result != "True" {
		t.Error("result fail",result)
	}
}
func CheckBal(id int,amt float32) string {

	bal := ""
	if id == 10002 {
		bal = "2000"
	}
	return bal
}

func GetBal()  {
	
}