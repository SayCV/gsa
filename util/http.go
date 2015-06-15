// Copyright (c) 2015 QDevor. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package util

import (
  `io/ioutil`
	`net/http`
	//`strings`
)

func ResponseToString(response *http.Response) (string, error) {
	content, err := ResponseToByteArray(response)
	return string(content), err
}

func ResponseToByteArray(response *http.Response) ([]byte, error) {
	body := response.Body
	return ioutil.ReadAll(body)
}