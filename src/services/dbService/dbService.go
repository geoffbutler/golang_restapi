package dbService

import (
	"database/sql"
	"fmt"
)

const driverName = "postgres"
const connectionString = "dbname=test user=postgres password=sql host=localhost port=5432"

// CheckDbConnectionStatus returns the status of the database connection
func CheckDbConnectionStatus() bool {
	fmt.Println("dbService: CheckDbConnectionStatus: BEGIN")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("getDbStatus: RECOVER")
		}
	}()

	db, connErr := sql.Open(driverName, connectionString)
	if connErr != nil {
		fmt.Println("dbService: CheckDbConnectionStatus: CONN ERROR")
		panic(connErr)
	}

	row := db.QueryRow("SELECT count(*) FROM test")

	var count int
	testErr := row.Scan(&count)
	if testErr != nil {
		fmt.Println("dbService: CheckDbConnectionStatus: TEST ERROR")
		panic(testErr)
	}

	fmt.Println("dbService: CheckDbConnectionStatus: END")
	return true
}
