// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	`fmt`
	//`strconv`
	`regexp`
	`strings`
	`github.com/SayCV/gsa/log`
)

func (quotes *Quotes) tencentParser(body []string) *Quotes {
	
	var regex *regexp.Regexp
	stockCnt := 0
	
	//log.Debug(fmt.Sprintf("Get body [%d] is ", len(body), body))
	for _, line := range body { 	
  	// log.Debug(fmt.Sprintf("Get line [%d] is [%s]", i, line))
  	// log.Debug("new line2")
  	
  	if !strings.Contains(line, `~`) {
  	  // log.Debug(`Lost error data.`)
  	  break
  	}
  	// log.Debug("new line3")
  	if strings.Contains(line, `pv_none_match`) {
  	  // log.Debug(`Get stock code none match.`)
  	  break
  	}
  	stockCnt++
  }
	quotes.stocks = make([]Stock, stockCnt)
	
	regex = regexp.MustCompile(`\"(.*)\"`)
	for i, line := range body { 	
  	log.Debug(fmt.Sprintf("Get line [%d] is [%s]", i, line))
  	// log.Debug("new line2")
  	
  	if !strings.Contains(line, `~`) {
  	  // log.Debug(`Lost error data.`)
  	  break
  	}
  	// log.Debug("new line3")
  	if strings.Contains(line, `pv_none_match`) {
  	  // log.Debug(`Get stock code none match.`)
  	  continue //break
  	}
  	// log.Debug("new line4")
  	matches := regex.FindStringSubmatch(string(line))
  	if len(matches) < 1 { break }
  	// log.Debug(fmt.Sprintf("Get regex [%d] is [%s]", len(matches), matches[1]))
  	matchesArray := strings.Split(matches[1], `~`)
  	// log.Debug(fmt.Sprintf("Get array [%d] is [%s]", len(matchesArray), matchesArray))
  	if len(matchesArray) < 44 {
			panic(`Unable to parse ` + string(i))
		}
  	// log.Debug("new line24")
    quotes.stocks[i].Name =                   matchesArray[1]
    quotes.stocks[i].Code =                   matchesArray[2]
    quotes.stocks[i].LastPrice = 							matchesArray[3]
    quotes.stocks[i].PrevPrice = 							matchesArray[4]
    quotes.stocks[i].OpenPrice = 							matchesArray[5]
    quotes.stocks[i].Volume = 								matchesArray[6]
    quotes.stocks[i].Timestamp = 							matchesArray[30]
    quotes.stocks[i].ChangePrice = 						matchesArray[31]
    quotes.stocks[i].ChangePricePct = 	      matchesArray[32]
    quotes.stocks[i].HighPrice = 							matchesArray[33]
    quotes.stocks[i].LowPrice = 							matchesArray[34]
    quotes.stocks[i].Amount = 								matchesArray[37]
    quotes.stocks[i].Swing = 									matchesArray[43]
    
    quotes.stocks[i].Advancing = 							quotes.stocks[i].ChangePrice > `0.00`
	}
	// log.Debug("new line000")
	return quotes
}


