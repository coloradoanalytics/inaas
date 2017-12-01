package main

import (
	"github.com/coloradoanalytics/go-iota-wrapper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	//NodeAddress         = "http://192.168.2.52:14680"     //address of the IRI
	//QuoteDuration       = 30                              //how long a quote is good for in seconds
	JavascriptISOString = "2006-01-02T15:04:05.999Z07:00" //for formatting time strings in a way that javascript likes
)

var (
	params AppParams
)

func main() {
	params = readParams()

	node := goiw.NewClient(params.NodeAddress)

	if !nodeIsReady(node) {
		log.Println("Node is not accessible or is not in sync")
		return
	}

	jobs := make(map[string]*Job)

	router := mux.NewRouter()

	router.HandleFunc("/quote", makeQuoteHandler(node, jobs))
	router.HandleFunc("/submit", makeSubmitHandler(node, jobs))
	router.HandleFunc("/status", makeStatusHandler(node, jobs))

	log.Fatal(http.ListenAndServe(":8080", router))

}

/////////////////// curl commands for testing /////////////////////
//curl http://localhost:8080/quote | python3 -m json.tool
//curl http://localhost:8080/submit -X POST -H 'Content-Type: application/json' -d '{"id": "8e9b2eba-4df9-4797-b0fc-a9a237203df6"}' | python3 -m json.tool
//curl http://localhost:8080/status -X POST -H 'Content-Type: application/json' -d '{"id": "8e9b2eba-4df9-4797-b0fc-a9a237203df6"}' | python3 -m json.tool
///////////////////////////////////////////////////////////////////

//
//
//
//////////////////////////////////////////////////////////////////////////////////////////////////
//Global Variables
//////////////////////////////////////////////////////////////////////////////////////////////////
// var client = iota.NewClient("http://192.168.2.52:14680")
// var quoteIDs = make(map[int]Quote)
// var txIDs = make(map[int]Transaction)
// var feePercentage = 0.01
// var averageTXTime int

//////////////////////////////////////////////////////////////////////////////////////////////////
//Handle calls.
//////////////////////////////////////////////////////////////////////////////////////////////////
// func mainHandler(w http.ResponseWriter, r *http.Request) {
// 	//Allow cross domain AJAX requests
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 	//Tell the browser we are sending JSON
// 	w.Header().Set("Content-Type", "application/json")

// 	action := r.URL.Query().Get("action")

// 	switch action {
// 	case "quote":
// 		amount := r.URL.Query().Get("amount")
// 		if amount != "" {
// 			intAmount, err := strconv.Atoi(amount)
// 			if err != nil {
// 				panic(err)
// 			}
// 			quote := getQuote(intAmount)
// 			fmt.Println(quote)
// 			json.NewEncoder(w).Encode(quote)
// 		} else {
// 			//Call to error handler with amount error.
// 		}
// 	case "transact":
// 		//call to transaction function with variables.
// 	case "status":
// 		txid := r.URL.Query().Get("txid")
// 		fmt.Println(txid)
// 		//Get status of transaction from IRI

// 	}
// }
