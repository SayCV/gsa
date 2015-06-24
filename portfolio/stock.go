// Copyright (c) 2015 by QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	//`github.com/michaeldv/termbox-go`
	//`regexp`
	//`strings`
	//`github.com/SayCV/gsa/log`
)

const MAX_DEALING_BLOCK = 5

type DealingBillQuote struct {
  Price             string
	Quantity					string
}

type DealingBill struct {
  Buy             *DealingBillQuote
	Sell            *DealingBillQuote
}

type Stock struct {
	Code				string
	Symbol			string
	Name				string
	MarketCap   string
	Board				string
	Industry		string
	
	PrevPrice				string
	OpenPrice				string
	LastPrice				string
	HighPrice				string
	LowPrice				string
	ChangePrice			string
	ChangePricePct	string
	LastVolume			string
	
	Volume					string
	Amount					string
	Swing						string
	TurnoverRatio		string
	VolumeRatio			string
	
	Dealing   [MAX_DEALING_BLOCK]DealingBill
	
	// milliseconds
	Timestamp				string
	
	Advancing       bool            // True when change is >= $0.

  AvgPrice        string
	PeRatio         string
	Dividend        string         // d: dividend.
	DividendYield   string         // y: dividend yield.
	
	errors          string         // Error(s), if any.
}

func (stock *Stock) GetError() string {
  return stock.errors
}

func (stock *Stock) SetError(error string) {
  stock.errors = error
}