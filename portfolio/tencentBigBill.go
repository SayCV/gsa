// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	`bytes`
	`fmt`
	`io/ioutil`
	`net/http`
	//`reflect`
	//`strconv`
	`strings`
  `github.com/SayCV/gsa/log`
)

const tencentBigBillQuotesURL = `http://stock.gtimg.cn/data/index.php?appn=radar&t=%s&d=09001515`

// Quotes stores relevant pointers as well as the array of stock quotes for
// the tickers we are tracking.
type BigBillQuotes struct {
	market  *ZhcnMarket  // Pointer to Market.
	profile *Profile // Pointer to Profile.
	stocks  []Stock  // Array of stock quote data.
	errors  string   // Error string if any.
}

// Sets the initial values and returns new Quotes struct.
func NewBigBillQuotes(market *ZhcnMarket, profile *Profile) *Quotes {
	return &BigBillQuotes{
		market:  market,
		profile: profile,
		errors:  ``,
	}
}

func (quotes *BigBillQuotes) GetProfile() *Profile {
  return quotes.profile
}

func (quotes *BigBillQuotes) GetStocks() []Stock {
  return quotes.stocks
}

func (quotes *BigBillQuotes) GetQueryCode(code string ) string {
  code = strings.ToLower(code)
  if strings.HasPrefix(code, `0`) || strings.HasPrefix(code, `3`) {
    return "sz" + code
  } else if strings.HasPrefix(code, `IF`) || strings.HasPrefix(code, `sh`) || strings.HasPrefix(code, `sz`) {
    return code
  } else {
    return "sh" + code
  }
}

// Fetch the latest stock quotes and parse raw fetched data into array of
// []Stock structs.
func (quotes *BigBillQuotes) Fetch() (self *BigBillQuotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		defer func() {
			if err := recover(); err != nil {
				quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
			}
		}()
    
    code := make([]string, len(quotes.profile.Tickers))
    for i, ticker := range quotes.profile.Tickers {
      code[i] = quotes.GetQueryCode(ticker)
    }
    
		url := fmt.Sprintf(zhcnQuotesURL, strings.Join(code, `,`))
		log.Debug(url)
		
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
    
		var stockJsonArray []string
    stockJsonArray = strings.Split(string(body), `;`)
	
	  // body = market.isMarketOpen(body)
    quotes.parse(stockJsonArray)
	}

	return quotes
}

// Ok returns two values: 1) boolean indicating whether the error has occured,
// and 2) the error text itself.
func (quotes *BigBillQuotes) Ok() (bool, string) {
	return quotes.errors == ``, quotes.errors
}


