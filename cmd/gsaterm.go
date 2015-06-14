// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package main

import (
	`time`
	"os"
	`github.com/michaeldv/termbox-go`
	`github.com/SayCV/gsa/log`
	`github.com/SayCV/gsa/portfolio`
	`github.com/SayCV/gsa/term`
	//`github.com/SayCV/gsa/util`
)

const help = `gsa v0.1.0 -- Copyright (c) 2015 by QDevor. All Rights Reserved.
NO WARRANTIES OF ANY KIND WHATSOEVER. SEE THE LICENSE FILE FOR DETAILS.

<u>Command</u>    <u>Description                                </u>
   +       Add stocks to the list.
   -       Remove stocks from the list.
   ?       Display this help screen.
   g       Group stocks by advancing/declining issues.
   o       Change column sort order.
   p       Pause market data and stock updates.
   q       Quit mop.
  esc      Ditto.
   b       BigBill.

Enter comma-delimited list of stock tickers when prompted.

<r> Press any key to continue </r>
`

//-----------------------------------------------------------------------------
func mainLoop(screen *term.Screen, profile *portfolio.Profile) {
	var lineEditor *term.LineEditor
	var columnEditor *term.ColumnEditor  

	keyboardQueue := make(chan termbox.Event)
	timestampQueue := time.NewTicker(1 * time.Second)
	quotesQueue := time.NewTicker(5 * time.Second)
	marketQueue := time.NewTicker(12 * time.Second)
	showingHelp := false
	paused := false

	go func() {
		for {
			keyboardQueue <- termbox.PollEvent()
		}
	}()

	market := portfolio.NewMarket()
	quotes := portfolio.NewQuotes(market, profile)
	screen.Draw(market, quotes)

loop:
	for {
		select {
		case event := <-keyboardQueue:
			switch event.Type {
			case termbox.EventKey:
				if lineEditor == nil && columnEditor == nil && !showingHelp {
					if event.Key == termbox.KeyEsc || event.Ch == 'q' || event.Ch == 'Q' {
						break loop
					} else if event.Ch == '+' || event.Ch == '-' {
						lineEditor = term.NewLineEditor(screen, quotes)
						lineEditor.Prompt(event.Ch)
					} else if event.Ch == 'o' || event.Ch == 'O' {
						columnEditor = term.NewColumnEditor(screen, quotes)
					} else if event.Ch == 'g' || event.Ch == 'G' {
						if profile.Regroup() == nil {
							screen.Draw(quotes)
						}
					} else if event.Ch == 'p' || event.Ch == 'P' {
						paused = !paused
						screen.Pause(paused).Draw(time.Now())
					} else if event.Ch == '?' || event.Ch == 'h' || event.Ch == 'H' {
						showingHelp = true
						screen.Clear().Draw(help)
					} else if event.Ch == 'b' || event.Ch == 'B' {
						showingHelp = true
						screen.Clear().Draw(`Coming soon`)
					}
				} else if lineEditor != nil {
					if done := lineEditor.Handle(event); done {
						lineEditor = nil
					}
				} else if columnEditor != nil {
					if done := columnEditor.Handle(event); done {
						columnEditor = nil
					}
				} else if showingHelp {
					showingHelp = false
					screen.Clear().Draw(market, quotes)
				}
			case termbox.EventResize:
				screen.Resize()
				if !showingHelp {
					screen.Draw(market, quotes)
				} else {
					screen.Draw(help)
				}
			}

		case <-timestampQueue.C:
			if !showingHelp && !paused {
				screen.Draw(time.Now())
			}

		case <-quotesQueue.C:
			if !showingHelp && !paused {
				screen.Draw(quotes)
			}

		case <-marketQueue.C:
			if !showingHelp && !paused {
				screen.Draw(market)
			}
		}
	}
}

//-----------------------------------------------------------------------------
func main() {
  
  //os.Setenv(`LOG_LEVEL`, `3`)
  os.Setenv(`LOG_TO_STDERR`, `true`)
  os.Setenv(`LOG_FILE_LOCATION`, `./`)
  log.Init()
  
  // flag.Parse()
	p, err := os.Getwd()  
  if err != nil {  
    log.Debug("Current Dir: ", err)  
  } else {  
    log.Debug("Current Dir: ", p)  
  }
  
  //util.Auxtest()
  //return
  
	screen := term.NewScreen()
	defer screen.Close()

	profile := portfolio.NewProfile()
	mainLoop(screen, profile)
	
	log.Flush()
}
