package inaas

import (
	"math/rand"
)

//////////////////////////////////////////////////////////////////////////////////////////////////
//Generate a random ID and verify it is not used.
//////////////////////////////////////////////////////////////////////////////////////////////////
func generateID(idType string) int{

	id := 100000 + rand.Intn(999999-100000)
	
	switch(idType){
		case "quote":
			if _, ok:= quoteIDs[id]; ok {         
				id = generateID(idType)
			}
		case "tx":
			if _, ok:= txIDs[id]; ok {         
				id = generateID(idType)
			}
	}
	return id
}