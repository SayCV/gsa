// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	`fmt`
	`strconv`
	`regexp`
	`strings`
	`github.com/SayCV/gsa/log`
)

func (quotes *Quotes) tencentParser(body []string) *Quotes {
	
	var regex *regexp.Regexp
	
	quotes.stocks = make([]Stock, len(body))
	
	regex = regexp.MustCompile(`\"(.*)\"`)
	for i, line := range body {
  	log.Debug(fmt.Sprintf("Get line [%d] is [%s]", i, line))
  	
  	matches := regex.FindStringSubmatch(string(line))
  	log.Debug("Get regex: ", len(matches), matches)
  	matchesArray := strings.Split(matches[1], `~`)
  	if len(matchesArray) < 44 {
			panic(`Unable to parse ` + string(i))
		}
  	
  	name := matchesArray[1]
    code := matchesArray[2]
    lastPrice, _ := 							strconv.ParseFloat(matchesArray[3], 32)
    prevPrice, _ := 							strconv.ParseFloat(matchesArray[4], 32)
    openPrice, _ := 							strconv.ParseFloat(matchesArray[5], 32)
    volume, _ := 									strconv.ParseUint(matchesArray[6], 10, 64)
    timestamp, _ := 							strconv.ParseUint(matchesArray[30], 10, 64)
    changePrice, _ := 						strconv.ParseFloat(matchesArray[31], 32)
    changePricePct, _ := 	        strconv.ParseFloat(matchesArray[32], 32)
    highPrice, _ := 							strconv.ParseFloat(matchesArray[33], 32)
    lowPrice, _ := 								strconv.ParseFloat(matchesArray[34], 32)
    amount, _ := 									strconv.ParseFloat(matchesArray[37], 32)
    swing, _ := 									strconv.ParseFloat(matchesArray[43], 32)
    
    quotes.stocks[i].name = name
    quotes.stocks[i].code = code
    quotes.stocks[i].lastPrice = 							float32(lastPrice)
    quotes.stocks[i].prevPrice = 							float32(prevPrice)
    quotes.stocks[i].openPrice = 							float32(openPrice)
    quotes.stocks[i].volume = 								volume
    quotes.stocks[i].timestamp = 							timestamp
    quotes.stocks[i].changePrice = 						float32(changePrice)
    quotes.stocks[i].changePricePct = 	      float32(changePricePct)
    quotes.stocks[i].highPrice = 							float32(highPrice)
    quotes.stocks[i].lowPrice = 							float32(lowPrice)
    quotes.stocks[i].amount = 								float32(amount)
    quotes.stocks[i].swing = 									float32(swing)
	}
	return quotes
}


