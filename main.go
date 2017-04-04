package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-oci8"
)

func main() {
	dsn := os.Getenv("GO_OCI8_CONNECT_STRING")
	db, err := sql.Open("oci8", dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	packageName := "Test"
	rows, err := db.Query("SELECT * FROM TABLE(INTG_PKG.GetP_packages_by_packagename(:1))", packageName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	fmt.Println(tableData)
	if err != nil {
		fmt.Println(err)
	}
}
