// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package util

import (
	`strconv`
	`strings`
)

func RuneToAscii(r rune) string {
    if r < 128 {
        return string(r)
    } else {
        return "\\u" + strconv.FormatInt(int64(r), 16)
    }
}

// The same exact method is used to sort by $Change and Change%. In both cases
// we sort by the value of Change% so that multiple $0.00s get sorted proferly.
func ChangeToFloat32(str string) float32 {
	trimmed := strings.Replace(strings.Trim(str, ` %`), `$`, ``, 1)
	value, _ := strconv.ParseFloat(trimmed, 32)
	return float32(value)
}

func Float32ToString(input_num float32) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', 2, 32)
}

func Float32ToStringM(input_num float32, bit int) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', bit, 32)
}