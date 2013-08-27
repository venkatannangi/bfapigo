package jsonparse

import (
	"encoding/json"
	"fmt"
	)
	

//market type definitions
type MaketType struct {
	MarketType string
	MarketCount int32
}

type MarketTypeResponse struct {	
	Result []MaketType
}

//competitions definitions
type Competition struct {
	Id string
	Name string
}

type CompetitionCount struct {
	Competition Competition
	MarketCount int32
}

type CompetitionsResponse struct {
	Result []CompetitionCount
}


type MarketBook struct {
	MarketId string
	IsMarketDataDelayed bool
	Status string
	BetDelay int32
	BspReconciled bool
	Complete bool
	Inplay bool
	NumberOfWinners int32
	NumberOfRunners int32
	NumberOfActiveRunners int32
	TotalMatched float32
	TotalAvailable float32
	CrossMatching bool
	
}

type MarketBookResponse struct {
	Result []MarketBook
}

func ParseMarketTypeResponse(response string) []MarketTypeResponse{
    var jsonBlob = []byte(response)	
	var responses []MarketTypeResponse
	
	err := json.Unmarshal(jsonBlob, &responses)
	
	if err != nil {
		fmt.Println("error:" ,err)
	}
	//fmt.Printf("%+v", responses)
    
    return 	responses
}


func ParseCompetitionsResponse(response string) []CompetitionsResponse{
    var jsonBlob = []byte(response)	
	var responses []CompetitionsResponse
	
	err := json.Unmarshal(jsonBlob, &responses)
	
	if err != nil {
		fmt.Println("error:" ,err)
	}
	//fmt.Printf("%+v", responses)
    
    return 	responses
}
