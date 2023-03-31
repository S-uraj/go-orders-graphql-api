package database

import (
	"fmt"

	"github.com/S-uraj/go-orders-graphql-api/graph/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// a variable to store database connection
var DBInstance *gorm.DB

// Var for error handling
var err error

// the db connection string
var CONNECTION_STRING string = "root:1020@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"

// connecting to the db
func ConnectDB() {
	// pass the db connection string
	ConnectionURI := CONNECTION_STRING
	// check for db connection
	DBInstance, err = gorm.Open("mysql", ConnectionURI)
	if err != nil {
		fmt.Println(err)
		// if the connection was unsuccessful
		panic("Database connection attempt was unsuccessful.....")
	} else {
		// if the connection was successful
		fmt.Println("Database Connected successfully.....")
	}
	// log all database operations performed by this connection
	DBInstance.LogMode(true)
}

func CreateDB() {
	// Create a database
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS Blog_Posts")
	// make the database available to this connection
	DBInstance.Exec("USE Blog_Posts")
}

func MigrateDB() {
	// migrate and sync the model to create a db table
	DBInstance.AutoMigrate(&model.Post{}, model.Comment{})
	//DBInstance.Model(&model.Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")

	fmt.Println("Database migration completed....")
}
