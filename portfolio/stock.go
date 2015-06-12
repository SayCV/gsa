// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	//`github.com/michaeldv/termbox-go`
	//`regexp`
	//`strings`
)


type Stock struct {
	code				string
	symbol			string
	name				string
	board				string
	industry		string
	
	prevPrice				float32
	openPrice				float32
	lastPrice				float32
	highPrice				float32
	lowPrice				float32
	changePrice									float32
	changePricePercentage				float32
	lastVolume			uint64
	
	volume					uint64
	amount					float32
	swing						float32
	turnoverRatio		float32
	volumeRatio			float32

	buyPrice							float32
	buyQuantity								uint64
	sellPrice							float32
	sellQuantity							uint64
	secondBuyPrice				float32
	secondBuyQuantity					uint64
	secondSellPrice				float32
	secondSellQuantity					uint64
	thirdBuyPrice					float32
	thirdBuyQuantity					uint64
	thirdSellPrice				float32
	thirdSellQuantity					uint64
	fourBuyPrice					float32
	fourBuyQuantity					uint64
	fourSellPrice				float32
	fourSellQuantity					uint64
	fiveBuyPrice					float32
	fiveBuyQuantity					uint64
	fiveSellPrice				float32
	fiveSellQuantity					uint64
	// milliseconds
	timestamp				uint64
	
	errors    string         // Error(s), if any.
}


