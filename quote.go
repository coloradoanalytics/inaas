package inaas

import (
	"math"
)

type Quote struct{
	ID		int `json:"ID"`
	Amount	int `json:"Amount"`
	Expires	int `json:"Expires"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Generate quote for proof of work fee
//////////////////////////////////////////////////////////////////////////////////////////////////
func getQuote(amount int) Quote{
	//Generate ID and add to ID slice
	newID := generateID("quote")
	//Do math for amount, either static or percentage
	fee := float64(float64(amount) * float64(feePercentage))
	rounded := math.Ceil(fee)
	//Calculate expiration, static or variables
	expirationTime := 10
	
	quote := Quote{ID: newID, Amount: int(rounded), Expires: expirationTime}
	quoteIDs[newID] = quote

	return quote
}