package main

import (
	"fmt"
	"os"

	_ "gopkg.in/rana/ora.v4"
	ora "gopkg.in/rana/ora.v4"
)

func main() {
	dsn := os.Getenv("OCI8_TEST_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}

	defer env.Close()
	defer srv.Close()
	defer ses.Close()

	stmtProcCall, err := ses.Prep("CALL PROC1(:1)")
	defer stmtProcCall.Close()
	if err != nil {
		panic(err)
	}
	procRset := &ora.Rset{}
	_, err = stmtProcCall.Exe(procRset)
	if err != nil {
		panic(err)
	}
	if procRset.IsOpen() {
		fmt.Println("before looping through rows")
		for procRset.Next() {
			fmt.Println("rows")
			fmt.Println(procRset.Row)
			//fmt.Println(procRset)
		}
		fmt.Println("after looping through rows")
		if err := procRset.Err(); err != nil {
			panic(err)
		}
		//fmt.Println(procRset.Len())
	}
}
