package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	packageName string
	expectedErr error
}{
	{
		packageName: "Get Bid",
		expectedErr: nil,
	},
}

func TestGetFunctionData(t *testing.T) {
	for _, c := range cases {
		env, srv, ses, _ := createIntegrationSession()

		defer env.Close()
		defer srv.Close()
		defer ses.Close()

		err := getFunctionData(ses, c.packageName)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
	}
}

func TestGetCursorData(t *testing.T) {
	env, srv, ses, _ := createLibrarianSession()

	defer env.Close()
	defer srv.Close()
	defer ses.Close()

	err := getCursorData(ses)
	if err != nil {
		t.Errorf("Expected err to be nil but it was %q", err)
	}

}
