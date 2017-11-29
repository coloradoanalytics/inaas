package main

import (
	"github.com/coloradoanalytics/go-iota-wrapper"
	"log"
)

//see if node is responding and synchronized
func nodeIsReady(node *goiw.Client) bool {
	info, err := node.GetNodeInfo()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	//node is synchronized if the milestone indices are equal
	return info.LatestMilestoneIndex == info.LatestSolidSubtangleMilestoneIndex
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//Check to see if node is responding!
//////////////////////////////////////////////////////////////////////////////////////////////////
// func nodeInit() bool {
// 	var initialized bool
// 	nodeInfo, err := client.GetNodeInfo()
// 	if err != nil {
// 		log.Fatal(err)
// 		initialized = false
// 	} else {
// 		log.Print(fmt.Sprintf("%+v", nodeInfo))
// 		initialized = true
// 	}

// 	return initialized
// }
