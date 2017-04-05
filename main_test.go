package main

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
)

var cases = []struct {
	dsn         string
	expectedErr error
}{
	{
		dsn:         "",
		expectedErr: errors.New("empty dsn"),
	},
}

func TestGetRecords(t *testing.T) {
	for _, c := range cases {
		oracle := &Oracle{}
		db, _ := sql.Open("oci8", c.dsn)
		oracle.db = db
		err := oracle.getRecords()
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}
	}
}
