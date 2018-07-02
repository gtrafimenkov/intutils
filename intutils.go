// SPDX-License-Identifier: MIT-0

package intutils

// AbsInt64 returns absolute value of an integer of int64 type
func AbsInt64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}
