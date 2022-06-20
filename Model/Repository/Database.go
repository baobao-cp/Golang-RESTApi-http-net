package Repository

import (
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	// "mysql", "<user>:<password>@tcp(127.0.0.1:3306)/<database-name>"
	driver := "mysql"
	user := "root"
	password := "12345"
	host := "localhost"
	port := "3306"
	databaseName := "golang"

	var err error
	db, err = sql.Open(driver, user+":"+password+"@tcp("+host+":"+port+")/"+databaseName)
	if err != nil {
		color.Red("Error: %s", err.Error())
		panic(err.Error())
	}
	color.Green("Database connection established successfully :  " + time.Now().String())

}

func closeDatabase() error {
	err := db.Close()
	if err != nil {
		return err
	}
	color.Yellow("Database connection closed successfully")
	return nil
}

func CreteTable(db *sql.DB) {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL
	);
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table created successfully")

}
