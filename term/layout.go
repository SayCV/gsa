// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package term

import (
	`bytes`
	`fmt`
	`reflect`
	`regexp`
	`strconv`
	`strings`
	`text/template`
	`time`
	`github.com/SayCV/gsa/log`
	//`github.com/SayCV/gsa/util`
	`github.com/SayCV/gsa/portfolio`
)

// Column describes formatting rules for individual column within the list
// of stock quotes.
type Column struct {
	width     int                 // Column width.
	name      string              // The name of the field in the Stock struct.
	title     string              // Column title to display in the header.
	formatter func(string) string // Optional function to format the contents of the column.
}

// Layout is used to format and display all the collected data, i.e. market
// updates and the list of stock quotes.
type Layout struct {
	columns        []Column           // List of stock quotes columns.
	sorter         *Sorter            // Pointer to sorting receiver.
	regex          *regexp.Regexp     // Pointer to regular expression to align decimal points.
	marketTemplate *template.Template // Pointer to template to format market data.
	quotesTemplate *template.Template // Pointer to template to format the list of stock quotes.
}

// Creates the layout and assigns the default values that stay unchanged.
func NewLayout() *Layout {
	layout := &Layout{}
	layout.columns = []Column{
		{-7, `Code`, `Ticker`, nil},
		{10, `LastPrice`, `Last`, currency},
		{10, `ChangePrice`, `Change`, currency},
		{10, `ChangePricePct`, `Change%`, last},
		{10, `OpenPrice`, `Open`, currency},
		{10, `LowPrice`, `Low`, currency},
		{10, `HighPrice`, `High`, currency},
		//{10, `Low52`, `52w Low`, currency},
		//{10, `High52`, `52w High`, currency},
		{11, `Volume`, `Volume`, nil},
		{11, `AvgPrice`, `AvgPrice`, nil},
		{9, `PeRatio`, `P/E`, blank},
		{9, `Dividend`, `Dividend`, zero},
		{9, `DividendYield`, `Yield`, percent},
		{11, `Amount`, `Amount`, currency},
	}
	layout.regex = regexp.MustCompile(`(\.\d+)[BMK]?$`)
	layout.marketTemplate = buildMarketTemplate()
	layout.quotesTemplate = buildQuotesTemplate()

	return layout
}

// Market merges given market data structure with the market template and
// returns formatted string that includes highlighting markup.
func (layout *Layout) Market(market *portfolio.ZhcnMarket) string {
	if ok, err := market.Ok(); !ok { // If there was an error fetching market data...
		return err // then simply return the error string.
	}

	highlight(
	  market.ShangHai, market.ShenZhen,
	  market.SmallPlate, market.GrowthEnterprise,
	  market.ShangHaiFund, market.HuShen300)
	  //market.Dow, market.Sp500, market.Nasdaq, market.HongKong)
	
	buffer := new(bytes.Buffer)
	layout.marketTemplate.Execute(buffer, market)

	return buffer.String()
}

// Quotes uses quotes template to format timestamp, stock quotes header,
// and the list of given stock quotes. It returns formatted string with
// all the necessary markup.
func (layout *Layout) Quotes(quotes *portfolio.Quotes) string {
	if ok, err := quotes.Ok(); !ok { // If there was an error fetching stock quotes...
		return err // then simply return the error string.
	}

	vars := struct {
		Now    string  // Current timestamp.
		Header string  // Formatted header line.
		Stocks []portfolio.Stock // List of formatted stock quotes.
	}{
		time.Now().Format(`| 3:04:05pm`),
		layout.Header(quotes.GetProfile()),
		layout.prettify(quotes),
	}

	buffer := new(bytes.Buffer)
	layout.quotesTemplate.Execute(buffer, vars)

	return buffer.String()
}

// Header iterates over column titles and formats the header line. The
// formatting includes placing an arrow next to the sorted column title.
// When the column editor is active it knows how to highlight currently
// selected column title.
func (layout *Layout) Header(profile *portfolio.Profile) string {
	str, selectedColumn := ``, profile.GetSelectedColumn()

	for i, col := range layout.columns {
		arrow := arrowFor(i, profile)
		if i != selectedColumn {
			str += fmt.Sprintf(`%*s`, col.width, arrow+col.title)
		} else {
			str += fmt.Sprintf(`<r>%*s</r>`, col.width, arrow+col.title)
		}
	}

	return `<u>` + str + `</u>`
}

// TotalColumns is the utility method for the column editor that returns
// total number of columns.
func (layout *Layout) TotalColumns() int {
	return len(layout.columns)
}

//-----------------------------------------------------------------------------
func (layout *Layout) prettify(quotes *portfolio.Quotes) []portfolio.Stock {
	pretty := make([]portfolio.Stock, len(quotes.GetStocks()))
	log.Debug("len(quotes.GetStocks()) is ", len(quotes.GetStocks()))
	//
	// Iterate over the list of stocks and properly format all its columns.
	//
	for i, stock := range quotes.GetStocks() {
		pretty[i].Advancing = stock.Advancing
		//
		// Iterate over the list of stock columns. For each column name:
		// - Get current column value.
		// - If the column has the formatter method then call it.
		// - Set the column value padding it to the given width.
		//
		for _, column := range layout.columns {
			// ex. value = stock.Change
			
			psGet := reflect.ValueOf(&stock) // pointer to struct - addressable
      sGet := psGet.Elem() // struct
      log.Debug("sGet is ", sGet)
      if sGet.Kind() != reflect.Struct {
        log.Debug(`sGet is Invalid`)
        return nil
      }
      // exported field
      fGet := sGet.FieldByName(column.name)
      log.Debug(`fGet is `, fGet)
      if !fGet.IsValid() {
        log.Debug(`fGet is Invalid`)
        return nil
      }
        // A Value can be changed only if it is 
        // addressable and was not obtained by 
        // the use of unexported struct fields.
        // if f.CanSet() {
          // change value of N
          // if f.Kind() == reflect.Int {
            // if !f.OverflowInt(x) {
            // f.SetInt(x)
            // }
            
          // }
        // }
      // value := reflect.ValueOf(&stock).Elem().FieldByName(column.name).String()
      var valueGet string
      switch fGet.Kind() {
      case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        valueGet = strconv.FormatInt(fGet.Int(), 10)
      case reflect.Float32:
        valueGet = strconv.FormatFloat(fGet.Float(), 'f', 2, 32)
      case reflect.String:
        valueGet = fGet.String()
      // etc...
      }
			log.Debug("valueGet is ", valueGet)
  		if column.formatter != nil {
  		  // ex. value = currency(valueGet)
  			// valueGet = column.formatter(valueGet)
  		}
    	// ex. pretty[i].Change = layout.pad(valueGet, 10)

    	//reflect.ValueOf(&pretty[i]).Elem().FieldByName(column.name).SetString(layout.pad(valueGet, column.width))
    	psSet := reflect.ValueOf(&pretty[i]) // pointer to struct - addressable
      sSet := psSet.Elem() // struct
      log.Debug(`sSet is `, sSet)
      if sSet.Kind() != reflect.Struct {
         log.Debug(`sSet is Invalid`)
         return nil
      }
      fSet := sSet.FieldByName(column.name)
      log.Debug(`fSet is `, fSet)
      if !fSet.IsValid() {
        log.Debug(`fSet is Invalid`)
        return nil
      }
    	
    	if !fSet.CanSet() {
    	  log.Debug(`fSet not CanSet`)
        return nil
    	}
    	// fSet.SetString(layout.pad(valueGet, column.width))
    	valueGet = layout.pad(valueGet, column.width)
    	log.Debug("valueGet by pad is ", valueGet)
    	switch fSet.Kind() {
      case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        valueSet, _ := strconv.ParseInt(valueGet, 10, 64)
        log.Debug(`valueSet [int] is `, valueSet)
        fSet.SetInt(valueSet)
      case reflect.Float32, reflect.Float64:
        valueSet, _ := strconv.ParseFloat(valueGet, 32)
        log.Debug(`valueSet [float] is `, valueSet)
        fSet.SetFloat(valueSet)
        log.Debug(`fSet [final] is `, strconv.FormatFloat(fSet.Float(), 'f', 2, 32))
      case reflect.String:
        log.Debug(`valueSet [string] is `, valueGet)
        fSet.SetString(valueGet)
      // etc...
      }
		}
	}

	profile := quotes.GetProfile()
	if layout.sorter == nil { // Initialize sorter on first invocation.
		layout.sorter = NewSorter(profile)
	}
	layout.sorter.SortByCurrentColumn(pretty)
	//
	// Group stocks by advancing/declining unless sorted by Chanage or Change%
	// in which case the grouping has been done already.
	//
	if profile.Grouped && (profile.SortColumn < 2 || profile.SortColumn > 3) {
		pretty = group(pretty)
	}

	return pretty
}

//-----------------------------------------------------------------------------
func (layout *Layout) pad(str string, width int) string {
	match := layout.regex.FindStringSubmatch(str)
	if len(match) > 0 {
		switch len(match[1]) {
		case 2:
			str = strings.Replace(str, match[1], match[1]+`0`, 1)
		case 4, 5:
			str = strings.Replace(str, match[1], match[1][0:3], 1)
		}
	}

	return fmt.Sprintf(`%*s`, width, str)
}

//-----------------------------------------------------------------------------
func buildMarketTemplate() *template.Template {
	markup := `{{if .ShangHai.advancing}}<red>SH   {{end}}</> {{.ShangHai.change}} ({{.ShangHai.percent}}) at  {{.ShangHai.latest}} <yellow>SZ   </> {{.ShenZhen.change}} ({{.ShenZhen.percent}}) at {{.ShenZhen.latest}}
<yellow>SPI  </> {{.SmallPlate.change}} ({{.SmallPlate.percent}}) at {{.SmallPlate.latest}} <yellow>GEI  </> {{.GrowthEnterprise.change}} ({{.GrowthEnterprise.percent}}) at {{.GrowthEnterprise.latest}}
<yellow>SHFI </> {{.ShangHaiFund.change}} ({{.ShangHaiFund.percent}}) at  {{.ShangHaiFund.latest}} <yellow>HS300</> {{.HuShen300.change}} ({{.HuShen300.percent}}) at {{.HuShen300.latest}})`

	return template.Must(template.New(`market`).Parse(markup))
}

//-----------------------------------------------------------------------------
func buildQuotesTemplate() *template.Template {
	markup := `<right><white>{{.Now}}</></right>



{{.Header}}
{{range.Stocks}}{{if .Advancing}}<red>{{end}}{{.Code}}{{.LastPrice}}{{.ChangePrice}}{{.ChangePricePct}}{{.OpenPrice}}{{.LowPrice}}{{.HighPrice}}{{/*.Low52*/}}{{/*.High52*/}}{{.Volume}}{{.AvgPrice}}{{.PeRatio}}{{.Dividend}}{{.DividendYield}}{{.Amount}}</>
{{end}}`

	return template.Must(template.New(`quotes`).Parse(markup))
}

//-----------------------------------------------------------------------------
func highlight(collections ...map[string]string) {
	for _, collection := range collections {
		if collection[`change`][0:1] != `-` {
			collection[`change`] = `<green>` + collection[`change`] + `</>`
		}
	}
}

//-----------------------------------------------------------------------------
func group(stocks []portfolio.Stock) []portfolio.Stock {
	grouped := make([]portfolio.Stock, len(stocks))
	current := 0

	for _, stock := range stocks {
		if stock.Advancing {
			grouped[current] = stock
			current++
		}
	}
	for _, stock := range stocks {
		if !stock.Advancing {
			grouped[current] = stock
			current++
		}
	}

	return grouped
}

//-----------------------------------------------------------------------------
func arrowFor(column int, profile *portfolio.Profile) string {
	if column == profile.SortColumn {
		if profile.Ascending {
			return string('\U00002191')
		}
		return string('\U00002193')
	}
	return ``
}

//-----------------------------------------------------------------------------
func blank(str string) string {
	if len(str) == 3 && str[0:3] == `N/A` {
		return `-`
	}

	return str
}

//-----------------------------------------------------------------------------
func zero(str string) string {
	if str == `0.00` {
		return `-`
	}

	return currency(str)
}

//-----------------------------------------------------------------------------
func last(str string) string {
	if len(str) >= 6 && str[0:6] == `N/A - ` {
		return str[6:]
	}

	return percent(str)
}

//-----------------------------------------------------------------------------
func currency(str string) string {
	if str == `N/A` {
		return `-`
	}
	if sign := str[0:1]; sign == `+` || sign == `-` {
		return sign + `$` + str[1:]
	}

	return `$` + str
}

// Returns percent value truncated at 2 decimal points.
//-----------------------------------------------------------------------------
func percent(str string) string {
	if str == `N/A` {
		return `-`
	}

	split := strings.Split(str, ".")
	if len(split) == 2 {
		digits := len(split[1])
		if digits > 2 {
			digits = 2
		}
		str = split[0] + "." + split[1][0:digits]
	}
	if str[len(str)-1] != '%' {
		str += `%`
	}
	return str
}
