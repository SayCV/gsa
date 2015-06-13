// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
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

type DealingQuote struct {
  Price             float32
	Quantity					uint64
}

type Dealing struct {
  Buy             *DealingQuote
	Sell            *DealingQuote
}

type Stock struct {
	Code				string
	Symbol			string
	Name				string
	MarketCap   string
	Board				string
	Industry		string
	
	PrevPrice				float32
	OpenPrice				float32
	LastPrice				float32
	HighPrice				float32
	LowPrice				float32
	ChangePrice						float32
	ChangePricePct				float32
	LastVolume			uint64
	
	Volume					uint64
	Amount					float32
	Swing						float32
	TurnoverRatio		float32
	VolumeRatio			float32

	// buyPrice							float32
	// buyQuantity								uint64
	// sellPrice							float32
	// sellQuantity							uint64
	// secondBuyPrice				float32
	// secondBuyQuantity					uint64
	// secondSellPrice				float32
	// secondSellQuantity					uint64
	// thirdBuyPrice					float32
	// thirdBuyQuantity					uint64
	// thirdSellPrice				float32
	// thirdSellQuantity					uint64
	// fourBuyPrice					float32
	// fourBuyQuantity					uint64
	// fourSellPrice				float32
	// fourSellQuantity					uint64
	// fiveBuyPrice					float32
	// fiveBuyQuantity					uint64
	// fiveSellPrice				float32
	// fiveSellQuantity					uint64
	
	Dealing   [MAX_DEALING_BLOCK]Dealing
	
	// milliseconds
	Timestamp				uint64
	
	Advancing       bool            // True when change is >= $0.

  AvgPrice        float32
	PeRatio         float32
	Dividend        float32         // d: dividend.
	DividendYield   float32         // y: dividend yield.
	
	errors          string         // Error(s), if any.
}

func (stock *Stock) GetError() string {
  return stock.errors
}

func (stock *Stock) SetError(error string) {
  stock.errors = error
}