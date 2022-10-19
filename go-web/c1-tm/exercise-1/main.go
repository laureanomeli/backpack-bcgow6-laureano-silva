package main

import (
	"encoding/json"
	"log"
	"os"
)

type transaction struct {
	Id       int
	Code     string
	Currency string
	Amount   float64
	Emitter  string
	Receiver string
	Date     string
}

func main() {

	firstTransaction := transaction{
		Id:       0,
		Code:     "a",
		Currency: "usd",
		Amount:   2000,
		Emitter:  "lalala",
		Receiver: "lelele",
		Date:     "03/10/2022",
	}
	transactions := []transaction{firstTransaction}

	transactions = append(transactions, transaction{
		Id:       1,
		Code:     "b",
		Currency: "usd",
		Amount:   3000,
		Emitter:  "lalala",
		Receiver: "lelele",
		Date:     "03/10/2022",
	},
		transaction{
			Id:       2,
			Code:     "b",
			Currency: "usd",
			Amount:   3300,
			Emitter:  "lalala",
			Receiver: "lelele",
			Date:     "03/10/2022",
		})

	jsonData, err := json.Marshal(transactions)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./transactions.json", jsonData, 666)
}
