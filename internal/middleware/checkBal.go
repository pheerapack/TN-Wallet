package middleware

type Wallet struct {
	WalletID int
	Bal int
}

func CheckBal(wallid int, bal int,amt int) string {
	if amt<bal {
		return "True"
	}
	return "False"
}


