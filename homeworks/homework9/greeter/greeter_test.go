package greeter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreeter(t *testing.T) {
	type testCase struct {
		name           string
		hour           int
		expectedResult string
	}

	tests := []testCase{
		{name: "Charly", hour: 6, expectedResult: "Good morning Charly!"},
		{name: "Charly", hour: 9, expectedResult: "Good morning Charly!"},
		{name: "Charly", hour: 11, expectedResult: "Good morning Charly!"},
		{name: "James", hour: 12, expectedResult: "Hello James!"},
		{name: "James", hour: 15, expectedResult: "Hello James!"},
		{name: "James", hour: 17, expectedResult: "Hello James!"},
		{name: "Stan", hour: 18, expectedResult: "Good evening Stan!"},
		{name: "Stan", hour: 21, expectedResult: "Good evening Stan!"},
		{name: "Stan", hour: 23, expectedResult: "Good evening Stan!"},
		{name: "Tom", hour: 0, expectedResult: "Good night Tom!"},
		{name: "Tom", hour: 3, expectedResult: "Good night Tom!"},
		{name: "Tom", hour: 5, expectedResult: "Good night Tom!"},
		{name: "Danny", hour: 99, expectedResult: "Hello Danny!"},
		{name: "danny", hour: 5, expectedResult: "Good night Danny!"},
		{name: "", hour: 11, expectedResult: "Good morning !"},
		{name: " Bobby ", hour: 11, expectedResult: "Good morning Bobby!"},
		{name: "Bobby ", hour: 11, expectedResult: "Good morning Bobby!"},
		{name: " Bobby", hour: 11, expectedResult: "Good morning Bobby!"},
		{name: "Bobby", hour: -11, expectedResult: "Hello Bobby!"},
	}

	for _, tc := range tests {
		testName := fmt.Sprintf("%v__request__%v", tc.name, tc.hour)
		t.Run(testName, func(t *testing.T) {
			actualResult := Greet(tc.name, tc.hour)
			assert.Equal(t, tc.expectedResult, actualResult, "Expected result: %v, Actual result: %v", tc.expectedResult, actualResult)
		})
	}

}
