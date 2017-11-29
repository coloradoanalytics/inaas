package main

import (
	"encoding/json"
	"github.com/coloradoanalytics/go-iota-wrapper"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

func makeQuoteHandler(node *goiw.Client, jobs map[string]*Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("quote request received")

		//create a new job
		job := &Job{
			ID:                uuid.NewV4().String(),
			Status:            "quoted",
			Fee:               10000,
			QuoteExpireTime:   time.Now().Add(QuoteDuration * time.Second),
			JobSubmissionChan: make(chan JobSubmission),
		}

		jobs[job.ID] = job

		quote := Quote{
			ID:      job.ID,
			Fee:     job.Fee,
			Expires: job.QuoteExpireTime.Format(JavascriptISOString),
		}

		data, err := json.Marshal(quote)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		go job.waitForSubmission() //start up the waitForSubmission function in a new goroutine so that it can run its own timer
		log.Println("quote response sent")
	}
}

type Quote struct {
	ID      string `json:"id"`
	Fee     int    `json:"fee"`
	Expires string `json:"expires"`
}

// type Quote struct {
// 	ID      int `json:"ID"`
// 	Amount  int `json:"Amount"`
// 	Expires int `json:"Expires"`
// }

//////////////////////////////////////////////////////////////////////////////////////////////////
//Generate quote for proof of work fee
//////////////////////////////////////////////////////////////////////////////////////////////////
// func getQuote(amount int) Quote {
// 	//Generate ID and add to ID slice
// 	newID := generateID("quote")
// 	//Do math for amount, either static or percentage
// 	fee := float64(float64(amount) * float64(feePercentage))
// 	rounded := math.Ceil(fee)
// 	//Calculate expiration, static or variables
// 	expirationTime := 10

// 	quote := Quote{ID: newID, Amount: int(rounded), Expires: expirationTime}
// 	quoteIDs[newID] = quote

// 	return quote
// }
