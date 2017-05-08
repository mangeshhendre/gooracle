package main

import (
	"reflect"
	"testing"
)

var functionCases = []struct {
	packageName string
	expectedErr error
}{
	{
		packageName: "Get Bid",
		expectedErr: nil,
	},
}

var cursorCases = []struct {
	orderNumber int64
	expectedErr error
}{
	{
		orderNumber: int64(600016555),
		expectedErr: nil,
	},
}

func TestGetFunctionDataPackages(t *testing.T) {
	for _, c := range functionCases {
		err := getFunctionDataPackages(c.packageName)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
	}
}

func TestGetCursorDataDepartments(t *testing.T) {

	err := getCursorDataDepartments()
	if err != nil {
		t.Errorf("Expected err to be nil but it was %q", err)
	}

}

func TestGetCursorDataImageDetails(t *testing.T) {
	for _, c := range cursorCases {
		err := getCursorDataImageDetails(c.orderNumber)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
	}
}
