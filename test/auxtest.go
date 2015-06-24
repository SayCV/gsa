// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package test

import (
  `fmt`
  //`strings`
  //`unicode/utf8`
	`github.com/SayCV/gsa/log`
	`github.com/SayCV/gsa/portfolio/ref`
)

func Auxtest() {
  log.Info("----------AuxTest ...----------")
  
  // auxtest1()
  auxtest2()
  
  log.Info("----------AuxTest ...----------")
}

func auxtest1() {
  //log.Info("----------AuxTest ...----------")
  
  token := `600497 驰宏锌锗      19.60     -0.19     -0.96     19.73     19.10     19.86     781246                                           152106`
  log.Info("token is: ", token)
  //r := []rune(token)
  //s := []byte(token)
  //log.Info("rune(token) is: ", r)
  //log.Info("byte(token) is: ", s)
  //
  //log.Info("len(token) is: ", len(token))
  //log.Info("utf8.RuneCount(s) is: ", utf8.RuneCount(s))
  //log.Info("utf8.RuneCountInString(token) is: ", utf8.RuneCountInString(token))
  //
  //log.Info("string(r) is: ", string(r))
  //log.Info("string(s) is: ", string(s))
  //r1, _ := utf8.DecodeRune(s)
  //log.Info("utf8.DecodeRune(s) is ", r1)
  
  for i, char := range token {
    //r, n := utf8.DecodeRune(s)
    if i < 20 {
      log.Info(fmt.Sprintf("range token by byte [%d] is %c", i, char))
    }
  }
  
  //log.Info("----------AuxTest ...----------")
}

func auxtest2() {
  ref.Test()
}
