package middleware

import (
	"testing"
)

func TestGetCheckBalLess(t *testing.T) {
	result := CheckBal(10002,2000,1000)

	if result != "True" {
		t.Error("result fail",result)
	}
}