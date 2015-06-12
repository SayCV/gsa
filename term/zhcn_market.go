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
	//`encoding/json`
	`github.com/SayCV/gsa/log`
	`github.com/SayCV/gsa/portfolio`
)

const zhcnMarketURL = `http://qt.gtimg.cn/q=%s`

var zhcnMarketCodes = map[string]string{
  `ShangHai`         : `sh000001`,
  `ShenZhen`         : `sz399001`,
  `GrowthEnterprise` : `sh000300`,
  `ShangHaiFund`     : `sz399006`,
  `ShenZhenFund`     : `sh000011`,
}

// Market stores current market information displayed in the top three lines of
// the screen. The market data is fetched and parsed from the HTML page above.
type zhcnMarket struct {
	IsClosed  bool              // True when China stcok markets are closed.
	
	ShangHai		map[string]string // Hash of S SECI Jones indicators.
	ShenZhen		map[string]string // Hash of SZ SECI indicators.
	GrowthEnterprise		map[string]string // Hash of GEI indicators.
	
	ShangHaiFund	map[string]string
	ShenZhenFund  map[string]string
	
	Dow       map[string]string // Hash of Dow Jones indicators.
	Nasdaq    map[string]string // Hash of NASDAQ indicators.
	Sp500     map[string]string // Hash of S&P 500 indicators.
	HongKong  map[string]string // Hash of HKHSI indicators.
  
  stocks  []Stock  // Array of stock quote data.
  
	regex     *regexp.Regexp // Regex to parse market data from HTML.
	errors    string         // Error(s), if any.
}

// Returns new initialized Market struct.
func NewMarket() *zhcnMarket {
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

func (market *zhcnMarket) getUrl() (string) {
  codes := fmt.Sprintf(`%s,%s`,
    zhcnMarketCodes[`ShangHai`],
    zhcnMarketCodes[`ShenZhen`])
  url := fmt.Sprintf(zhcnMarketURL, codes)
  
  log.Debug(url)

  return url
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
  _zhcnMarketURL := market.getUrl()
	response, err := http.Get(_zhcnMarketURL)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
  
  // log.Debug("Get body: ", string(body))
  
  var dataArray []string
  dataArray = strings.Split(string(body), `;`)
	
	// body = market.isMarketOpen(body)
	return market.extract(dataArray)
}

// Ok returns two values: 1) boolean indicating whether the error has occured,
// and 2) the error text itself.
func (market *zhcnMarket) Ok() (bool, string) {
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
func (market *zhcnMarket) extract(snippet []string) *zhcnMarket {
	// matches := market.regex.FindStringSubmatch(string(snippet))
	var matches []string
	
	portfolio.Stock

	if len(matches) < 44 {
		panic(`Unable to parse ` + zhcnMarketURL)
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

	return market
}
