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
	marketCap   string
	board				string
	industry		string
	
	prevPrice				float32
	openPrice				float32
	lastPrice				float32
	highPrice				float32
	lowPrice				float32
	changePrice						float32
	changePricePct				float32
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
	
	advancing       bool            // True when change is >= $0.

  avgPrice        float32
	peRatio         float32
	dividend        float32         // d: dividend.
	dividendYield   float32         // y: dividend yield.
	
	errors          string         // Error(s), if any.
}

func (stock *Stock) GetCode() string {
  return stock.code
}

func (stock *Stock) GetTicker() string {
  return stock.code
}

func (stock *Stock) GetSymbol() string {
  return stock.symbol
}

func (stock *Stock) GetName() string {
  return stock.name
}

func (stock *Stock) GetMarketCap() string {
  return stock.marketCap
}

func (stock *Stock) GetBoard() string {
  return stock.board
}

func (stock *Stock) GetIndustry() string {
  return stock.industry
}

func (stock *Stock) GetPrevPrice() float32 {
  return stock.prevPrice
}
	
func (stock *Stock) GetOpenPrice() float32 {
  return stock.openPrice
}

func (stock *Stock) GetLastPrice() float32 {
  return stock.lastPrice
}

func (stock *Stock) GetHighPrice() float32 {
  return stock.highPrice
}

func (stock *Stock) GetLowPrice() float32 {
  return stock.lowPrice
}

func (stock *Stock) GetChangePrice() float32 {
  return stock.changePrice
}

func (stock *Stock) GetChangePricePct() float32 {
  return stock.changePricePct
}

func (stock *Stock) GetLastVolume() uint64 {
  return stock.lastVolume
}

func (stock *Stock) GetVolume() uint64 {
  return stock.volume
}

func (stock *Stock) GetAmount() float32 {
  return stock.amount
}

func (stock *Stock) GetSwing() float32 {
  return stock.swing
}
	
func (stock *Stock) GetTurnoverRatio() float32 {
  return stock.turnoverRatio
}
	
func (stock *Stock) GetVolumeRatio() float32 {
  return stock.volumeRatio
}

	//buyPrice							float32
	//buyQuantity								uint64
	//sellPrice							float32
	//sellQuantity							uint64
	//secondBuyPrice				float32
	//secondBuyQuantity					uint64
	//secondSellPrice				float32
	//secondSellQuantity					uint64
	//thirdBuyPrice					float32
	//thirdBuyQuantity					uint64
	//thirdSellPrice				float32
	//thirdSellQuantity					uint64
	//fourBuyPrice					float32
	//fourBuyQuantity					uint64
	//fourSellPrice				float32
	//fourSellQuantity					uint64
	//fiveBuyPrice					float32
	//fiveBuyQuantity					uint64
	//fiveSellPrice				float32

func (stock *Stock) GetFiveSellQuantity() uint64 {
  return stock.fiveSellQuantity
}

func (stock *Stock) GetTimestamp() uint64 {
  return stock.timestamp
}

func (stock *Stock) GetAdvancing() bool {
  return stock.advancing
}

func (stock *Stock) GetAvgPrice() float32 {
  return stock.avgPrice
}

func (stock *Stock) GetPeRatio() float32 {
  return stock.peRatio
}

func (stock *Stock) GetDividend() float32 {
  return stock.dividend
}

func (stock *Stock) GetDividendYield() float32 {
  return stock.dividendYield
}

func (stock *Stock) GetErrors() string {
  return stock.errors
}


func (stock *Stock) SetCode(code string) {
  stock.code = code
}
	
func (stock *Stock) SetSymbol(symbol string) {
  stock.symbol = symbol
}

func (stock *Stock) SetName(name string) {
  stock.name = name
}

func (stock *Stock) SetMarketCap(marketCap string) {
  stock.marketCap = marketCap
}

func (stock *Stock) SetBoard(board string) {
  stock.board = board
}

func (stock *Stock) SetIndustry(industry string) {
  stock.industry = industry
}

func (stock *Stock) SetPrevPrice(prevPrice float32) {
  stock.prevPrice = prevPrice
}
	
func (stock *Stock) SetOpenPrice(openPrice float32) {
  stock.openPrice = openPrice
}

func (stock *Stock) SetLastPrice(lastPrice float32) {
  stock.lastPrice = lastPrice
}

func (stock *Stock) SetHighPrice(highPrice float32) {
  stock.highPrice = highPrice
}

func (stock *Stock) SetLowPrice(lowPrice float32) {
  stock.lowPrice = lowPrice
}

func (stock *Stock) SetChangePrice(changePrice float32) {
  stock.changePrice = changePrice
}

func (stock *Stock) SetChangePricePct(changePricePct float32) {
  stock.changePricePct = changePricePct
}

func (stock *Stock) SetLastVolume(lastVolume uint64) {
  stock.lastVolume = lastVolume
}

func (stock *Stock) SetVolume(volume uint64) {
  stock.volume = volume
}

func (stock *Stock) SetAmount(amount float32) {
  stock.amount = amount
}

func (stock *Stock) SetSwing(swing float32) {
  stock.swing = swing
}

func (stock *Stock) SetTurnoverRatio(turnoverRatio float32) {
  stock.turnoverRatio = turnoverRatio
}
	
func (stock *Stock) SetVolumeRatio(volumeRatio float32) {
  stock.volumeRatio = volumeRatio
}

	//buyPrice							float32
	//buyQuantity								uint64
	//sellPrice							float32
	//sellQuantity							uint64
	//secondBuyPrice				float32
	//secondBuyQuantity					uint64
	//secondSellPrice				float32
	//secondSellQuantity					uint64
	//thirdBuyPrice					float32
	//thirdBuyQuantity					uint64
	//thirdSellPrice				float32
	//thirdSellQuantity					uint64
	//fourBuyPrice					float32
	//fourBuyQuantity					uint64
	//fourSellPrice				float32
	//fourSellQuantity					uint64
	//fiveBuyPrice					float32
	//fiveBuyQuantity					uint64
	//fiveSellPrice				float32

func (stock *Stock) SetFiveSellQuantity(fiveSellQuantity uint64) {
  stock.fiveSellQuantity = fiveSellQuantity
}

func (stock *Stock) SetTimestamp(timestamp uint64) {
  stock.timestamp = timestamp
}

func (stock *Stock) SetAdvancing(advancing bool) {
  stock.advancing = advancing
}

func (stock *Stock) SetAvgPrice(avgPrice float32) {
  stock.avgPrice = avgPrice
}

func (stock *Stock) SetPeRatio(peRatio float32) {
  stock.peRatio = peRatio
}

func (stock *Stock) SetDividend(dividend float32) {
  stock.dividend = dividend
}

func (stock *Stock) SetDividendYield(dividendYield float32) {
  stock.dividendYield = dividendYield
}

func (stock *Stock) SetErrors(errors string) {
  stock.errors = errors
}
