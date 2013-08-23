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
	Jsonrpc string
	Result []MaketType
	Id int32
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
	Jsonrpc string
	Result []CompetitionCount
	Id int32
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
