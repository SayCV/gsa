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
	stockCnt := 0
	
	log.Debug(fmt.Sprintf("Get body [%d] is ", len(body), body))
	for i, line := range body { 	
  	log.Debug(fmt.Sprintf("Get line [%d] is [%s]", i, line))
  	log.Debug("new line2")
  	
  	if !strings.Contains(line, `~`) {
  	  log.Debug(`Lost error data.`)
  	  break
  	}
  	log.Debug("new line3")
  	if strings.Contains(line, `pv_none_match`) {
  	  log.Debug(`Get stock code none match.`)
  	  break
  	}
  	stockCnt++
  }
	quotes.stocks = make([]Stock, stockCnt)
	
	regex = regexp.MustCompile(`\"(.*)\"`)
	for i, line := range body { 	
  	log.Debug(fmt.Sprintf("Get line [%d] is [%s]", i, line))
  	log.Debug("new line2")
  	
  	if !strings.Contains(line, `~`) {
  	  log.Debug(`Lost error data.`)
  	  break
  	}
  	log.Debug("new line3")
  	if strings.Contains(line, `pv_none_match`) {
  	  log.Debug(`Get stock code none match.`)
  	  continue //break
  	}
  	log.Debug("new line4")
  	matches := regex.FindStringSubmatch(string(line))
  	if len(matches) < 1 { break }
  	log.Debug(fmt.Sprintf("Get regex [%d] is [%s]", len(matches), matches[1]))
  	matchesArray := strings.Split(matches[1], `~`)
  	log.Debug(fmt.Sprintf("Get array [%d] is [%s]", len(matchesArray), matchesArray))
  	if len(matchesArray) < 44 {
			panic(`Unable to parse ` + string(i))
		}
  	log.Debug("new line24")
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
	log.Debug("new line000")
	return quotes
}


