// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package term

import (
	`sort`
	`strconv`
	`strings`
	//`github.com/SayCV/gsa/util`
	`github.com/SayCV/gsa/portfolio`
)

// Sorter gets called to sort stock quotes by one of the columns. The
// setup is rather lengthy; there should probably be more concise way
// that uses reflection and avoids hardcoding the column names.
type Sorter struct {
	profile *portfolio.Profile // Pointer to where we store sort column and order.
}

type sortable []portfolio.Stock

func (list sortable) Len() int      { return len(list) }
func (list sortable) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

type byTickerAsc struct{ sortable }
type byLastTradeAsc struct{ sortable }
type byChangeAsc struct{ sortable }
type byChangePctAsc struct{ sortable }
type byOpenAsc struct{ sortable }
type byLowAsc struct{ sortable }
type byHighAsc struct{ sortable }

type byVolumeAsc struct{ sortable }
type byAvgPriceAsc struct{ sortable }
type byPeRatioAsc struct{ sortable }
type byDividendAsc struct{ sortable }
type byYieldAsc struct{ sortable }
type byMarketCapAsc struct{ sortable }

type byTickerDesc struct{ sortable }
type byLastTradeDesc struct{ sortable }
type byChangeDesc struct{ sortable }
type byChangePctDesc struct{ sortable }
type byOpenDesc struct{ sortable }
type byLowDesc struct{ sortable }
type byHighDesc struct{ sortable }

type byVolumeDesc struct{ sortable }
type byAvgPriceDesc struct{ sortable }
type byPeRatioDesc struct{ sortable }
type byDividendDesc struct{ sortable }
type byYieldDesc struct{ sortable }
type byMarketCapDesc struct{ sortable }

func (list byTickerAsc) Less(i, j int) bool {
	return list.sortable[i].Code < list.sortable[j].Code
}
func (list byLastTradeAsc) Less(i, j int) bool {
	return list.sortable[i].LastPrice < list.sortable[j].LastPrice
}
func (list byChangeAsc) Less(i, j int) bool {
	return c(list.sortable[i].ChangePrice) < c(list.sortable[j].ChangePrice)
}
func (list byChangePctAsc) Less(i, j int) bool {
	return c(list.sortable[i].ChangePricePct) < c(list.sortable[j].ChangePricePct)
}
func (list byOpenAsc) Less(i, j int) bool {
	return list.sortable[i].OpenPrice < list.sortable[j].OpenPrice
}
func (list byLowAsc) Less(i, j int) bool { return list.sortable[i].LowPrice < list.sortable[j].LowPrice }
func (list byHighAsc) Less(i, j int) bool {
	return list.sortable[i].HighPrice < list.sortable[j].HighPrice
}


func (list byVolumeAsc) Less(i, j int) bool {
	return list.sortable[i].Volume < list.sortable[j].Volume
}
func (list byAvgPriceAsc) Less(i, j int) bool {
	return list.sortable[i].AvgPrice < list.sortable[j].AvgPrice
}
func (list byPeRatioAsc) Less(i, j int) bool {
	return list.sortable[i].PeRatio < list.sortable[j].PeRatio
}
func (list byDividendAsc) Less(i, j int) bool {
	return list.sortable[i].Dividend < list.sortable[j].Dividend
}
func (list byYieldAsc) Less(i, j int) bool {
	return list.sortable[i].DividendYield < list.sortable[j].DividendYield
}
func (list byMarketCapAsc) Less(i, j int) bool {
	return m(list.sortable[i].MarketCap) < m(list.sortable[j].MarketCap)
}

func (list byTickerDesc) Less(i, j int) bool {
	return list.sortable[j].Code < list.sortable[i].Code
}
func (list byLastTradeDesc) Less(i, j int) bool {
	return list.sortable[j].LastPrice < list.sortable[i].LastPrice
}
func (list byChangeDesc) Less(i, j int) bool {
	return c(list.sortable[j].ChangePrice) < c(list.sortable[i].ChangePrice)
}
func (list byChangePctDesc) Less(i, j int) bool {
	return c(list.sortable[j].ChangePricePct) < c(list.sortable[i].ChangePricePct)
}
func (list byOpenDesc) Less(i, j int) bool {
	return list.sortable[j].OpenPrice < list.sortable[i].OpenPrice
}
func (list byLowDesc) Less(i, j int) bool { return list.sortable[j].LowPrice < list.sortable[i].LowPrice }
func (list byHighDesc) Less(i, j int) bool {
	return list.sortable[j].HighPrice < list.sortable[i].HighPrice
}



func (list byVolumeDesc) Less(i, j int) bool {
	return list.sortable[j].Volume < list.sortable[i].Volume
}
func (list byAvgPriceDesc) Less(i, j int) bool {
	return list.sortable[j].AvgPrice < list.sortable[i].AvgPrice
}
func (list byPeRatioDesc) Less(i, j int) bool {
	return list.sortable[j].PeRatio < list.sortable[i].PeRatio
}
func (list byDividendDesc) Less(i, j int) bool {
	return list.sortable[j].Dividend < list.sortable[i].Dividend
}
func (list byYieldDesc) Less(i, j int) bool {
	return list.sortable[j].DividendYield < list.sortable[i].DividendYield
}
func (list byMarketCapDesc) Less(i, j int) bool {
	return m(list.sortable[j].MarketCap) < m(list.sortable[i].MarketCap)
}

// Returns new Sorter struct.
func NewSorter(profile *portfolio.Profile) *Sorter {
	return &Sorter{
		profile: profile,
	}
}

// SortByCurrentColumn builds a list of sort interface based on current sort
// order, then calls sort.Sort to do the actual job.
func (sorter *Sorter) SortByCurrentColumn(stocks []portfolio.Stock) *Sorter {
	var interfaces []sort.Interface

	if sorter.profile.Ascending {
		interfaces = []sort.Interface{
			byTickerAsc{stocks},
			byLastTradeAsc{stocks},
			byChangeAsc{stocks},
			byChangePctAsc{stocks},
			byOpenAsc{stocks},
			byLowAsc{stocks},
			byHighAsc{stocks},
			byVolumeAsc{stocks},
			byAvgPriceAsc{stocks},
			byPeRatioAsc{stocks},
			byDividendAsc{stocks},
			byYieldAsc{stocks},
			byMarketCapAsc{stocks},
		}
	} else {
		interfaces = []sort.Interface{
			byTickerDesc{stocks},
			byLastTradeDesc{stocks},
			byChangeDesc{stocks},
			byChangePctDesc{stocks},
			byOpenDesc{stocks},
			byLowDesc{stocks},
			byHighDesc{stocks},
			byVolumeDesc{stocks},
			byAvgPriceDesc{stocks},
			byPeRatioDesc{stocks},
			byDividendDesc{stocks},
			byYieldDesc{stocks},
			byMarketCapDesc{stocks},
		}
	}

	sort.Sort(interfaces[sorter.profile.SortColumn])

	return sorter
}

// The same exact method is used to sort by $Change and Change%. In both cases
// we sort by the value of Change% so that multiple $0.00s get sorted proferly.
func c(str string) float32 {
	trimmed := strings.Replace(strings.Trim(str, ` %`), `$`, ``, 1)
	value, _ := strconv.ParseFloat(trimmed, 32)
	return float32(value)
}

// When sorting by the market value we must first convert 42B etc. notations
// to proper numeric values.
func m(str string) float32 {
	multiplier := 1.0

	switch str[len(str)-1 : len(str)] { // Check the last character.
	case `B`:
		multiplier = 1000000000.0
	case `M`:
		multiplier = 1000000.0
	case `K`:
		multiplier = 1000.0
	}

	trimmed := strings.Trim(str, ` $BMK`) // Get rid of non-numeric characters.
	value, _ := strconv.ParseFloat(trimmed, 32)

	return float32(value * multiplier)
}
