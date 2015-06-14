// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package portfolio

import (
	//`regexp`
	//`strings`
	//`github.com/SayCV/gsa/log`
)

type StockBigBill struct {
	Code				string
	Name				string
  Price       string
  Info        string
  Id		      string
	
	errors          string         // Error(s), if any.
}

func (sbb *StockBigBill) GetError() string {
  return sbb.errors
}

func (sbb *StockBigBill) SetError(error string) {
  sbb.errors = error
}