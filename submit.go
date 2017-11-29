package main

import (
	"encoding/json"
	"github.com/coloradoanalytics/go-iota-wrapper"
	"log"
	"net/http"
)

func makeSubmitHandler(node *goiw.Client, jobs map[string]*Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("submit request received")

		decoder := json.NewDecoder(r.Body)
		var s JobSubmission
		err := decoder.Decode(&s)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		response := JobSubmissionResponse{ID: s.ID}

		job, ok := jobs[s.ID]

		if ok {
			log.Println("id found")
			response.Status = "OK"
			if job.quoteIsValid() {
				job.JobSubmissionChan <- s //put the job submission from the client into the job's submission channel
			} else {
				response.Status = "quote expired"
			}
		} else {
			log.Println("id not found")
			response.Status = "Invalid ID"
		}

		data, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)

		log.Println("submit response sent")
	}
}

type JobSubmission struct {
	ID        string `json:"id"`        //job id assigned by server
	FeeTrytes string `json:"feeTrytes"` //trytes for fee transaction to pay the server
	JobTrytes string `json:"jobTrytes"` //trytes for job transaction
}

type JobSubmissionResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
