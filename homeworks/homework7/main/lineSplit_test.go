package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine(t *testing.T) {
	type testCase struct {
		request        string
		expectedResult string
	}

	tests := []testCase{
		{request: "hello world. with pleasure", expectedResult: "Hello world. With pleasure."},
		{request: "hello world. with pleasure.", expectedResult: "Hello world. With pleasure."},
		{request: "hello world. With pleasure", expectedResult: "Hello world. With pleasure."},
		{request: "Hello world. with pleasure", expectedResult: "Hello world. With pleasure."},
		{request: "hello world. with pleasure. super test", expectedResult: "Hello world. With pleasure. Super test."},
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("%v__convert_to__%v", tc.request, tc.expectedResult)
		t.Run(testName, func(t *testing.T) {
			actualResult := fixString(tc.request)
			assert.Equal(t, tc.expectedResult, actualResult, "Expected result: %v, Actual result: %v", tc.expectedResult, actualResult)
		})

	}

}
