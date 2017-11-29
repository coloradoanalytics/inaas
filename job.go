package main

import (
	"log"
	"time"
)

type Job struct {
	ID                string             //UUID to identify this quote/job when talking to the client
	Status            string             //current status of quote/job
	Fee               int                //price in IOTA quoted to the client to send a transaction
	QuoteExpireTime   time.Time          //time when the quote will expire
	JobSubmissionChan chan JobSubmission //channel used to send a job submission to the waitForSubmission function
	FeeTrytes         string             //trytes of signed transaction to pay the server, provided by client
	JobTrytes         string             //trytes of signed transaction server is being paid to send, provided by client
}

//after a quote has been sent, wait until an order is submitted with the quote ID or until the quote expiration time passes
func (j *Job) waitForSubmission() {
	select {
	case <-time.After(j.QuoteExpireTime.Sub(time.Now())):
		//quote timer expired
		j.expireQuote()
	case submission := <-j.JobSubmissionChan:
		//job received
		j.startJob(submission)
	}
}

func (j *Job) expireQuote() {
	log.Println("timer expired")
	j.Status = "expired"
}

func (j *Job) startJob(submission JobSubmission) {
	log.Println("starting job")
	j.Status = "processing"
}

func (j *Job) quoteIsValid() bool {
	return j.Status == "quoted"
}
