package util

import (
	"errors"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestStringToInt(t *testing.T) {
	testCase := []struct {
		name           string
		expectedInput  string
		expectedOutput int
	}{
		{
			name:           "error parser, number with space character",
			expectedInput:  " 0",
			expectedOutput: 0,
		},
		{
			name:           "error parser, another character",
			expectedInput:  "ss221~+",
			expectedOutput: 0,
		},
		{
			name:           "success parser int",
			expectedInput:  "0",
			expectedOutput: 0,
		},
		{
			name:           "big int",
			expectedInput:  "999999999999999999999999999999999999999999999999999999999999999",
			expectedOutput: 0,
		},
		{
			name:           "minus",
			expectedInput:  "-1",
			expectedOutput: -1,
		},
	}

	for _, c := range testCase {
		t.Run(c.name, func(t *testing.T) {
			var actualOutput int
			StringToInt(&actualOutput, c.expectedInput, 0)

			if actualOutput != c.expectedOutput {
				t.Errorf("parser int expected  %v but got %v", c.expectedOutput, actualOutput)
			}
		})
	}
}

func TestStringToTime(t *testing.T) {

	defaultValue := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)

	testCase := []struct {
		name           string
		expectedInput  string
		expectedOutput time.Time
	}{
		{
			name:           "error parser, empty param",
			expectedInput:  "",
			expectedOutput: defaultValue,
		},
		{
			name:           "error parser, invalid format",
			expectedInput:  "2020-19-04T07:57:50.103Z",
			expectedOutput: defaultValue,
		},
		{
			name:           "success parser time",
			expectedInput:  "1900-01-01T00:00:00.000Z",
			expectedOutput: defaultValue,
		},
	}

	for _, c := range testCase {
		t.Run(c.name, func(t *testing.T) {
			var actualOutput time.Time
			StringToTime(&actualOutput, c.expectedInput)

			if actualOutput != c.expectedOutput {
				t.Errorf("parser time expected %v but got %v", c.expectedOutput, actualOutput)
			}
		})
	}
}

func TestStringToObjectID(t *testing.T) {

	defaultValue, _ := primitive.ObjectIDFromHex("5fa8cbe94530a2bb653a851c")

	testCase := []struct {
		name           string
		expectedInput  string
		expectedOutput primitive.ObjectID
		expectedErr    error
	}{
		{
			name:          "error parser, empty param",
			expectedInput: "",
			expectedErr:   errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name:          "error parser, invalid format",
			expectedInput: "5fa8cbe94530a2bb653a85c",
			expectedErr:   errors.New("encoding/hex: odd length hex string"),
		},
		{
			name:           "success parser objectid",
			expectedInput:  "5fa8cbe94530a2bb653a851c",
			expectedOutput: defaultValue,
			expectedErr:    errors.New("ee"),
		},
	}

	for _, c := range testCase {
		t.Run(c.name, func(t *testing.T) {
			var actualOutput primitive.ObjectID
			actualErr := StringToObjectID(&actualOutput, c.expectedInput)
			if actualErr != nil && actualErr.Error() != c.expectedErr.Error() {
				t.Errorf("parser ObjectID expected %v but got %v", c.expectedErr, actualErr)
			} else {
				if actualOutput != c.expectedOutput {
					t.Errorf("parser ObjectID expected %v but got %v", c.expectedErr, actualErr)
				}
			}
		})
	}
}
