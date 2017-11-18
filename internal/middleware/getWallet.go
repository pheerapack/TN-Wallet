package middleware

import (
	"strconv"
	"net/http"
	"mux"
	"log"
)


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

func GetWallet (w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	myString := vars["walletID"]

	log.Println(myString)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(myString))


}

func queryWallet(id int) string {

	info := ""
	if id == 20001 {

		acc := Account{}
		acc.Mer()

		info = strconv.Itoa(acc.Id)+"|"+acc.Name+"|"+acc.Pin

	}
	return info
}
