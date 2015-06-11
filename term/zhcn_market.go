// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package term

import (
	`bytes`
	`fmt`
	`io/ioutil`
	`net/http`
	`regexp`
	`strings`
)

const zhcnMarketURL = `http://money.cnn.com/data/markets/`

// Market stores current market information displayed in the top three lines of
// the screen. The market data is fetched and parsed from the HTML page above.
type zhcnMarket struct {
	IsClosed  bool              // True when China stcok markets are closed.
	
	shIndex		map[string]string // Hash of S SECI Jones indicators.
	szIndex		map[string]string // Hash of SZ SECI indicators.
	geIndex		map[string]string // Hash of GEI indicators.
	
	shFIndex	map[string]string
	szFIndex  map[string]string
	
	Dow       map[string]string // Hash of Dow Jones indicators.
	Nasdaq    map[string]string // Hash of NASDAQ indicators.
	Sp500     map[string]string // Hash of S&P 500 indicators.
	HongKong  map[string]string // Hash of HKHSI indicators.
	
	regex     *regexp.Regexp // Regex to parse market data from HTML.
	errors    string         // Error(s), if any.
}

// Returns new initialized Market struct.
func zhcnNewMarket() *zhcnMarket {
	market := &zhcnMarket{}
	market.IsClosed = false
	
	market.ShangHai = make(map[string]string)
	market.ShenZhen = make(map[string]string)
	market.GrowthEnterprise = make(map[string]string)
	
	market.ShangHaiFund = make(map[string]string)
	market.ShenZhenFund = make(map[string]string)
	
	market.Dow = make(map[string]string)
	market.Nasdaq = make(map[string]string)
	market.Sp500 = make(map[string]string)
	market.HongKong = make(map[string]string)
	
	
	market.errors = ``

	const any = `\s*(?:.+?)`
	const price = `>([\d\.,]+)</span>`
	const percent = `>([\+\-]?[\d\.,]+%?)<`

	rules := []string{
		`>ShangHai<`, any, percent, any, price, any, percent, any,
		`>ShenZhen<`, any, percent, any, price, any, percent, any,
		`>GrowthEnterprise<`, any, percent, any, price, any, percent, any,
		`>ShangHaiFund<`, any, percent, any, price, any, percent, any,
		`>ShenZhenFund<`, any, percent, any, price, any, percent, any,
		`>Dow<`, any, percent, any, price, any, percent, any,
		`>Nasdaq<`, any, percent, any, price, any, percent, any,
		`">S&P<`, any, percent, any, price, any, percent, any,
		`>Hang Seng<`, any, percent, any, price, any, percent, any,
	}

	market.regex = regexp.MustCompile(strings.Join(rules, ``))

	return market
}

// Fetch downloads HTML page from the 'zhcnMarketURL', parses it, and stores resulting data
// in internal hashes. If download or data parsing fails Fetch populates 'market.errors'.
func (market *zhcnMarket) Fetch() (self *zhcnMarket) {
	self = market // <-- This ensures we return correct market after recover() from panic().
	defer func() {
		if err := recover(); err != nil {
			market.errors = fmt.Sprintf("Error fetching market data...\n%s", err)
		}
	}()

	response, err := http.Get(zhcnMarketURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	body = market.isMarketOpen(body)
	return market.extract(market.trim(body))
}

// Ok returns two values: 1) boolean indicating whether the error has occured,
// and 2) the error text itself.
func (market *Market) Ok() (bool, string) {
	return market.errors == ``, market.errors
}

//-----------------------------------------------------------------------------
func (market *zhcnMarket) isMarketOpen(body []byte) []byte {
	// TBD -- CNN page doesn't seem to have market open/close indicator.
	return body
}

//-----------------------------------------------------------------------------
func (market *zhcnMarket) trim(body []byte) []byte {
	start := bytes.Index(body, []byte(`Markets Overview`))
	finish := bytes.LastIndex(body, []byte(`Gainers`))
	snippet := bytes.Replace(body[start:finish], []byte{'\n'}, []byte{}, -1)
	snippet = bytes.Replace(snippet, []byte(`&amp;`), []byte{'&'}, -1)

	return snippet
}

//-----------------------------------------------------------------------------
func (market *zhcnMarket) extract(snippet []byte) *zhcnMarket {
	matches := market.regex.FindStringSubmatch(string(snippet))

	if len(matches) < 31 {
		panic(`Unable to parse ` + marketURL)
	}

	market.ShangHai[`change`] = matches[1]
	market.ShangHai[`latest`] = matches[2]
	market.ShangHai[`percent`] = matches[3]

	market.ShenZhen[`change`] = matches[4]
	market.ShenZhen[`latest`] = matches[5]
	market.ShenZhen[`percent`] = matches[6]

	market.GrowthEnterprise[`change`] = matches[7]
	market.GrowthEnterprise[`latest`] = matches[8]
	market.GrowthEnterprise[`percent`] = matches[9]
	
	market.ShangHaiFund[`change`] = matches[7]
	market.ShangHaiFund[`latest`] = matches[8]
	market.ShangHaiFund[`percent`] = matches[9]
	
	market.ShenZhenFund[`change`] = matches[7]
	market.ShenZhenFund[`latest`] = matches[8]
	market.ShenZhenFund[`percent`] = matches[9]
	
	market.Dow[`change`] = matches[1]
	market.Dow[`latest`] = matches[2]
	market.Dow[`percent`] = matches[3]

	market.Nasdaq[`change`] = matches[4]
	market.Nasdaq[`latest`] = matches[5]
	market.Nasdaq[`percent`] = matches[6]

	market.Sp500[`change`] = matches[7]
	market.Sp500[`latest`] = matches[8]
	market.Sp500[`percent`] = matches[9]

	market.HongKong[`change`] = matches[13]
	market.HongKong[`latest`] = matches[14]
	market.HongKong[`percent`] = matches[15]

	return market
}
