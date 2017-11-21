package main

import (
	"fmt"
	"github.com/coloradoanalytics/iota"
	"log"
)

func main() {
	client := iota.NewClient("http://192.168.2.52:14680")

	// nodeInfo, err := client.GetNodeInfo()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(fmt.Sprintf("%+v", nodeInfo))

	// neighbors, err := client.GetNeighbors()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(fmt.Sprintf("%+v", neighbors.Neighbors))

	tips, err := client.GetTips()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(fmt.Sprintf("%+v", tips.Hashes))

}
