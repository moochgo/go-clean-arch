package mysql

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	//Side effect GORM by Dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBs ...
type DBs struct {
	Customer *gorm.DB
}

var DBConnect = DBs{}

var db *gorm.DB

// Initialize Method: Connection Test to DB Server
func Initialize(dbName string) *gorm.DB {
	connSetting := "charset=utf8mb4&parseTime=True"
	connString := os.Getenv("DB_MYSQL_USER") + ":" + os.Getenv("DB_MYSQL_PASS") + "@(" + os.Getenv("DB_MYSQL_IP") + ")/" + dbName + "?" + connSetting
	log.Println(">", connString)

	db, err := gorm.Open(os.Getenv("DB_MYSQL"), connString)
	if err != nil {
		log.Println(fmt.Sprintf("[DB SRV] Error Connection Testing to DB - %v", err))
		return nil
	}

	log.Println(fmt.Sprintf("[DB SRV] Successful Connection Testing to DB: %v", connString))
	// db pool
	db.DB().SetMaxOpenConns(0)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetConnMaxLifetime(0)
	active, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	db.LogMode(active)
	return db
}

// CustomerConnection ...
func (s *DBs) CustomerConnection() {
	s.Customer = Initialize("persada_customers")
}
