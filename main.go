package inaas

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/coloradoanalytics/iota"
	"log"
)

//////////////////////////////////////////////////////////////////////////////////////////////////
//Global Variables
//////////////////////////////////////////////////////////////////////////////////////////////////
var client = iota.NewClient("http://192.168.2.52:14680")
var quoteIDs = make(map[int]Quote)
var txIDs = make(map[int]Transaction)
var feePercentage = 0.01
var averageTXTime int

func initAvgTime() {
	//Run a few 0 value txs and get the average time to complete.
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Check to see if node is responding!
//////////////////////////////////////////////////////////////////////////////////////////////////
func nodeInit() bool{
	var initialized bool
	nodeInfo, err := client.GetNodeInfo()
	if err != nil {
		log.Fatal(err)
		initialized = false
	} else {
		log.Print(fmt.Sprintf("%+v", nodeInfo))
		initialized = true
	}

	return initialized
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Handle calls.
//////////////////////////////////////////////////////////////////////////////////////////////////
func mainHandler(w http.ResponseWriter, r *http.Request){
	//Allow cross domain AJAX requests
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
    //Tell the browser we are sending JSON
	w.Header().Set("Content-Type", "application/json")
	
	action := r.URL.Query().Get("action")
	
	switch(action){
		case "quote":
			amount := r.URL.Query().Get("amount")
			if amount != ""{
				intAmount, err := strconv.Atoi(amount)
				if err != nil{
					panic(err)
				}
				quote := getQuote(intAmount)
				fmt.Println(quote)
				json.NewEncoder(w).Encode(quote)
			} else {
				//Call to error handler with amount error.
			}
		case "transact":
			//call to transaction function with variables.
		case "status":
			txid := r.URL.Query().Get("txid")
			fmt.Println(txid)
			//Get status of transaction from IRI
		
	}
}

func main() {
	if nodeInit() {
		http.HandleFunc("/", mainHandler)
		http.ListenAndServe(":443", nil)
	}
	// neighbors, err := client.GetNeighbors()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(fmt.Sprintf("%+v", neighbors.Neighbors))

	// tips, err := client.GetTips()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(fmt.Sprintf("%+v", tips.Hashes))
}