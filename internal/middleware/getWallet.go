package middleware

import "strconv"


type Account struct {
	Id int  `bson:"ID"`
	Name string	 `bson:"Name"`
	Pin	string
}

func (a *Account) Mer() {
	a.Id = 20001
	a.Name = "7 - 14"
	a.Pin = "1234"
}

func getWallet(id int) string {

	info := ""
	if id == 20001 {

		acc := Account{}
		acc.Mer()

		info = strconv.Itoa(acc.Id)+"|"+acc.Name+"|"+acc.Pin

	}
	return info
}
