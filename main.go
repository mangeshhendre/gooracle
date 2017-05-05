package main

import (
	"fmt"
	"os"

	_ "gopkg.in/rana/ora.v4"
	ora "gopkg.in/rana/ora.v4"
)

func main() {
	env, srv, ses, err := createIntegrationSession()
	if err != nil {
		fmt.Println(err)
	}
	defer env.Close()
	defer srv.Close()
	defer ses.Close()

	err = getFunctionData(ses, "Get Bid")
	if err != nil {
		fmt.Println(err)
	}
}

func createIntegrationSession() (*ora.Env, *ora.Srv, *ora.Ses, error) {
	dsn := os.Getenv("GO_OCI8_INTG_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		return nil, nil, nil, err
	}

	return env, srv, ses, nil
}

func createLibrarianSession() (*ora.Env, *ora.Srv, *ora.Ses, error) {
	dsn := os.Getenv("GO_OCI8_LIB_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		return nil, nil, nil, err
	}

	return env, srv, ses, nil
}

func getCursorData(session *ora.Ses) error {
	//Prepare the query
	prepStatement, err := session.Prep("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)")
	if err != nil {
		return err
	}
	defer prepStatement.Close()

	//Retrieve the resultSet
	resultSet := &ora.Rset{}
	_, err = prepStatement.Exe(resultSet)
	if err != nil {
		return err
	}

	//Trying to print the values retured from ref_cursor
	fmt.Println("Cursor Test")
	rowCount := 0
	if resultSet.IsOpen() {
		//Print columns
		for _, v := range resultSet.Columns {
			fmt.Print(v.Name, " ")
		}
		fmt.Println()

		//Print rows
		for resultSet.Next() {
			for k := range resultSet.Columns {
				fmt.Print(resultSet.Row[k], " ")
			}
			rowCount++
			fmt.Println()
		}
		fmt.Println("Number of rows returned:", rowCount)
		//rset.Len() also doesn't seem to work nil pointer dereference
		//fmt.Println("Number of rows returned:", resultSet.Len())
	}

	return nil
}

func getFunctionData(session *ora.Ses, packageName string) error {
	//Call sql function to get list of packages
	prepStatement, err := session.Prep("SELECT * FROM TABLE(INTG_PKG.GetP_packages_by_packagename(:1))")
	if err != nil {
		return err
	}
	defer prepStatement.Close()

	//Execute the function
	rset, err := prepStatement.Qry(packageName)
	if err != nil {
		return err
	}

	//Print results
	fmt.Println("Function Test")
	rowCount := 0
	if rset.IsOpen() {

		//Print columns
		for _, v := range rset.Columns {
			fmt.Print(v.Name, " ")
		}
		fmt.Println()
		//Print rows
		for rset.Next() {
			for k := range rset.Columns {
				fmt.Print(rset.Row[k], " ")
			}
			rowCount++
			fmt.Println()
		}
		fmt.Println("Number of rows returned:", rowCount)
	}

	return nil
}
