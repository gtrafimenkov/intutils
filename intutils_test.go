// SPDX-License-Identifier: MIT-0

package intutils

import (
	"fmt"
	"testing"
)

func TestAbsInt64(t *testing.T) {
	type TestCase struct {
		input  int64
		output int64
	}

	tests := []TestCase{
		TestCase{0, 0},
		TestCase{10, 10},
		TestCase{-10, 10},
	}

	for _, test := range tests {
		output := AbsInt64(test.input)
		if output != test.output {
			t.Errorf("Expecting %v, got %v", test.output, output)
		}
	}
}

func ExampleAbsInt64() {
	fmt.Println(AbsInt64(-123))

	// Output: 123
}
