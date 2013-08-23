package main

import (
	"fmt"
	"com.betfair.api/ng/callapi"
	"com.betfair.api/ng/jsonparse"
	)
	
func main() {
    var applicationKey = "your application key" //palnning to pass as command line arguments soon
    var url = "https://api.betfair.com/exchange/betting/json-rpc/v1"
    var sessionToken = "your session token" //palnning to pass as command line arguments soon
    var marketTypePayload = "[{\"jsonrpc\":\"2.0\",\"method\":\"SportsAPING/v1.0/listMarketTypes\", \"params\": [{\"filter\":{}}],\"id\": 1}]"
    var competitionsPayload = "[{\"jsonrpc\":\"2.0\",\"method\":\"SportsAPING/v1.0/listCompetitions\", \"params\": [{\"filter\":{}}],\"id\": 1}]"	
	marketTypeResponse := callapi.SendRequest(url,applicationKey,sessionToken, marketTypePayload)
	competitionsResponse := callapi.SendRequest(url,applicationKey,sessionToken, competitionsPayload)
	
	fmt.Println("-------------Market Types-----------")
	fmt.Printf("%+v", jsonparse.ParseMarketTypeResponse(marketTypeResponse))
	fmt.Println("-------------Competitions Types-----------")
	fmt.Printf("%+v", jsonparse.ParseCompetitionsResponse(competitionsResponse))
}
