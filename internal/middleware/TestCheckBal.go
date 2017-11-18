package middleware

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mux"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func ResponseWithJSON(w http.ResponseWriter, code int) {

	w.Header().Set("datetime", time.Now().Format("2006-01-02 15:04:05+0700"))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-roundtrip", "")
	w.WriteHeader(code)

}

//Struct id,amt 10002 1,000
type WalletIDAmt struct {
	WalletID string
	Amount   float32
}

type WalletAccount struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	WalletID int           `json:"wallet_id" bson:"wallet_id"`
	Name     string        `json:"wallet_name" bson:"wallet_name"`
	Balance  float32       `json:"balance" bson:"balance"`
}

func CheckBal(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		vars := mux.Vars(r)
		wallets, _ := strconv.Atoi(vars["balance"])

		var account WalletAccount
		c := session.DB("wallets").C("balance")
		err := c.Find(bson.M{"wallet_id": wallets}).One(&account)
		if err != nil {
			ResponseWithJSON(w, http.StatusBadRequest)
		}
		ResponseWithJSON(w, http.StatusOK)

	}
}

func TestCheckBal(t testing.T) {

}
