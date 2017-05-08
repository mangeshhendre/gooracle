package main

import (
	"fmt"
	"os"

	_ "gopkg.in/rana/ora.v4"
	ora "gopkg.in/rana/ora.v4"
)

func main() {
	err := getCursorDataDepartments()
	if err != nil {
		fmt.Println(err)
	}

	err = getCursorDataImageDetails()
	if err != nil {
		fmt.Println(err)
	}
}

func getCursorDataImageDetails() error {
	dsn := os.Getenv("GO_OCI8_LIB_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}

	defer env.Close()
	defer srv.Close()
	defer ses.Close()
	//Prepare the query
	prepStatement, err := ses.Prep("CALL CONTENTSERVICE.RETRIEVEIMAGEDETAILLIST(:1,:2)")
	if err != nil {
		return err
	}
	defer prepStatement.Close()

	//Retrieve the resultSet
	resultSet := &ora.Rset{}
	_, err = prepStatement.Exe(int64(600016555), resultSet)
	if err != nil {
		return err
	}

	//Trying to print the values retured from ref_cursor
	fmt.Println("Cursor Test")
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
			fmt.Println()
		}
		fmt.Println("Number of rows returned:", resultSet.Len())
	}

	return nil
}

func getCursorDataDepartments() error {
	dsn := os.Getenv("GO_OCI8_LIB_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}

	defer env.Close()
	defer srv.Close()
	defer ses.Close()
	//Prepare the query
	prepStatement, err := ses.Prep("CALL CONTENTSERVICE.RETRIEVEDEPARTMENTS(:1)")
	if err != nil {
		return err
	}
	defer prepStatement.Close()

	//Retrieve the resultSet
	resultSet := &ora.Rset{}
	rowsEffected, err := prepStatement.Exe(resultSet)
	if err != nil {
		return err
	}
	fmt.Println("rowsEffected", rowsEffected)
	//Trying to print the values retured from ref_cursor
	fmt.Println("Cursor Test")
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
			fmt.Println()
		}
		fmt.Println("Number of rows returned:", resultSet.Len())
	}

	return nil
}

func getFunctionData(session *ora.Ses, packageName string) error {
	dsn := os.Getenv("GO_OCI8_INTG_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}

	defer env.Close()
	defer srv.Close()
	defer ses.Close()
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
			fmt.Println()
		}
		fmt.Println("Number of rows returned:", rset.Len())
	}

	return nil
}
