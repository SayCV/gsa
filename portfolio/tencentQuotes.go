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
	`strings`
  `github.com/SayCV/gsa/log`
)

// See http://www.gummy-stuff.org/Yahoo-stocks.htm
//
// const quotesURL = `http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=sl1c6k2oghjkva2r2rdyj3j1`
// c2: realtime change vs c1: change
// k2: realtime change vs p2: change
//
const zhcnQuotesURL = `http://qt.gtimg.cn/q=%s`

// Quotes stores relevant pointers as well as the array of stock quotes for
// the tickers we are tracking.
type Quotes struct {
	market  *ZhcnMarket  // Pointer to Market.
	profile *Profile // Pointer to Profile.
	stocks  []Stock  // Array of stock quote data.
	errors  string   // Error string if any.
}

// Sets the initial values and returns new Quotes struct.
func NewQuotes(market *ZhcnMarket, profile *Profile) *Quotes {
	return &Quotes{
		market:  market,
		profile: profile,
		errors:  ``,
	}
}

func (quotes *Quotes) GetProfile() *Profile {
  return quotes.profile
}

func (quotes *Quotes) GetStocks() []Stock {
  return quotes.stocks
}

// Fetch the latest stock quotes and parse raw fetched data into array of
// []Stock structs.
func (quotes *Quotes) Fetch() (self *Quotes) {
	self = quotes // <-- This ensures we return correct quotes after recover() from panic().
	if quotes.isReady() {
		defer func() {
			if err := recover(); err != nil {
				quotes.errors = fmt.Sprintf("\n\n\n\nError fetching stock quotes...\n%s", err)
			}
		}()
    
		url := fmt.Sprintf(zhcnQuotesURL, strings.Join(quotes.profile.Tickers, `,`))
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
func (quotes *Quotes) Ok() (bool, string) {
	return quotes.errors == ``, quotes.errors
}

// AddTickers saves the list of tickers and refreshes the stock data if new
// tickers have been added. The function gets called from the line editor
// when user adds new stock tickers.
func (quotes *Quotes) AddTickers(tickers []string) (added int, err error) {
	if added, err = quotes.profile.AddTickers(tickers); err == nil && added > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

// RemoveTickers saves the list of tickers and refreshes the stock data if some
// tickers have been removed. The function gets called from the line editor
// when user removes existing stock tickers.
func (quotes *Quotes) RemoveTickers(tickers []string) (removed int, err error) {
	if removed, err = quotes.profile.RemoveTickers(tickers); err == nil && removed > 0 {
		quotes.stocks = nil // Force fetch.
	}
	return
}

// isReady returns true if we haven't fetched the quotes yet *or* the stock
// market is still open and we might want to grab the latest quotes. In both
// cases we make sure the list of requested tickers is not empty.
func (quotes *Quotes) isReady() bool {
	return (quotes.stocks == nil || !quotes.market.IsClosed) && len(quotes.profile.Tickers) > 0
}

// Use reflection to parse and assign the quotes data fetched using the Yahoo
// market API.
func (quotes *Quotes) parse(body []string) *Quotes {
	return quotes.tencentParser(body)
}

//-----------------------------------------------------------------------------
func sanitize(body []byte) []byte {
	return bytes.Replace(bytes.TrimSpace(body), []byte{'"'}, []byte{}, -1)
}
