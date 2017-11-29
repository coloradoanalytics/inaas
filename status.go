package main

import (
	"encoding/json"
	"github.com/coloradoanalytics/go-iota-wrapper"
	"log"
	"net/http"
)

func makeStatusHandler(node *goiw.Client, jobs map[string]*Job) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("status request received")

		decoder := json.NewDecoder(r.Body)

		request := struct {
			ID string
		}{}

		err := decoder.Decode(&request)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		response := StatusResponse{ID: request.ID}

		job, ok := jobs[request.ID]

		if ok {
			log.Println("id found")
			response.Status = job.Status
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

		log.Println("status response sent")
	}
}

type StatusResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Check status of transaction.
//////////////////////////////////////////////////////////////////////////////////////////////////
// func getStatus(txid string){
// 	//Ask IRI about status of transaction
// 	//Update in map
// 	//return status.
// }
