// Copyright (c) 2013-2015 by Michael Dvorkin. All Rights Reserved.
// Use of this source code is governed by a MIT-style license that can
// be found in the LICENSE file.

package util

import (
	`strconv`
)

func Float32ToString(input_num float32) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', 2, 32)
}

func Float32ToStringM(input_num float32, bit int) string {
    // to convert a float number to a string
    return strconv.FormatFloat(float64(input_num), 'f', bit, 32)
}