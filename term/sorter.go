// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package term

import (
	`sort`
	`strconv`
	`strings`
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

func float32ToString(input_num float32) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', 2, 32)
}

func (list byTickerAsc) Less(i, j int) bool {
	return list.sortable[i].GetTicker() < list.sortable[j].GetTicker()
}
func (list byLastTradeAsc) Less(i, j int) bool {
	return list.sortable[i].GetLastVolume() < list.sortable[j].GetLastVolume()
}
func (list byChangeAsc) Less(i, j int) bool {
	return c(float32ToString(list.sortable[i].GetChangePrice())) < c(float32ToString(list.sortable[j].GetChangePrice()))
}
func (list byChangePctAsc) Less(i, j int) bool {
	return c(float32ToString(list.sortable[i].GetChangePricePct())) < c(float32ToString(list.sortable[j].GetChangePricePct()))
}
func (list byOpenAsc) Less(i, j int) bool {
	return list.sortable[i].GetOpenPrice() < list.sortable[j].GetOpenPrice()
}
func (list byLowAsc) Less(i, j int) bool { return list.sortable[i].GetLowPrice() < list.sortable[j].GetLowPrice() }
func (list byHighAsc) Less(i, j int) bool {
	return list.sortable[i].GetHighPrice() < list.sortable[j].GetHighPrice()
}


func (list byVolumeAsc) Less(i, j int) bool {
	return list.sortable[i].GetVolume() < list.sortable[j].GetVolume()
}
func (list byAvgPriceAsc) Less(i, j int) bool {
	return list.sortable[i].GetAvgPrice() < list.sortable[j].GetAvgPrice()
}
func (list byPeRatioAsc) Less(i, j int) bool {
	return list.sortable[i].GetPeRatio() < list.sortable[j].GetPeRatio()
}
func (list byDividendAsc) Less(i, j int) bool {
	return list.sortable[i].GetDividend() < list.sortable[j].GetDividend()
}
func (list byYieldAsc) Less(i, j int) bool {
	return list.sortable[i].GetDividendYield() < list.sortable[j].GetDividendYield()
}
func (list byMarketCapAsc) Less(i, j int) bool {
	return m(list.sortable[i].GetMarketCap()) < m(list.sortable[j].GetMarketCap())
}

func (list byTickerDesc) Less(i, j int) bool {
	return list.sortable[j].GetTicker() < list.sortable[i].GetTicker()
}
func (list byLastTradeDesc) Less(i, j int) bool {
	return list.sortable[j].GetLastVolume() < list.sortable[i].GetLastVolume()
}
func (list byChangeDesc) Less(i, j int) bool {
	return c(float32ToString(list.sortable[j].GetChangePrice())) < c(float32ToString(list.sortable[i].GetChangePrice()))
}
func (list byChangePctDesc) Less(i, j int) bool {
	return c(float32ToString(list.sortable[j].GetChangePricePct())) < c(float32ToString(list.sortable[i].GetChangePricePct()))
}
func (list byOpenDesc) Less(i, j int) bool {
	return list.sortable[j].GetOpenPrice() < list.sortable[i].GetOpenPrice()
}
func (list byLowDesc) Less(i, j int) bool { return list.sortable[j].GetLowPrice() < list.sortable[i].GetLowPrice() }
func (list byHighDesc) Less(i, j int) bool {
	return list.sortable[j].GetHighPrice() < list.sortable[i].GetHighPrice()
}



func (list byVolumeDesc) Less(i, j int) bool {
	return list.sortable[j].GetVolume() < list.sortable[i].GetVolume()
}
func (list byAvgPriceDesc) Less(i, j int) bool {
	return list.sortable[j].GetAvgPrice() < list.sortable[i].GetAvgPrice()
}
func (list byPeRatioDesc) Less(i, j int) bool {
	return list.sortable[j].GetPeRatio() < list.sortable[i].GetPeRatio()
}
func (list byDividendDesc) Less(i, j int) bool {
	return list.sortable[j].GetDividend() < list.sortable[i].GetDividend()
}
func (list byYieldDesc) Less(i, j int) bool {
	return list.sortable[j].GetDividendYield() < list.sortable[i].GetDividendYield()
}
func (list byMarketCapDesc) Less(i, j int) bool {
	return m(list.sortable[j].GetMarketCap()) < m(list.sortable[i].GetMarketCap())
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
