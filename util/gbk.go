// Copyright (c) 2015 QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package util

import (
  `bytes`
  `github.com/axgle/mahonia`
  `golang.org/x/text/encoding/simplifiedchinese`
  `golang.org/x/text/transform`
  `io/ioutil`
)

func GbkEncode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
    if err == nil {
        dst = string(data)
    }
    return
}

func GbkDecode(src string) (dst string) {
    data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
    if err == nil {
        dst = string(data)
    }
    return
}

func MahoniaEncode(src string) (dst string) {
    enc := mahonia.NewEncoder("gbk")
    return enc.ConvertString(src)
}

func MahoniaDecode(src string) (dst string) {
    dec := mahonia.NewDecoder("gbk")
    return dec.ConvertString(src)
}