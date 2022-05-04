package main

import (
	"database/sql"
	"fmt"
	"github.com/heptiolabs/healthcheck"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var UriDb = "postgres://user:passs@host:port/testDb"

func connectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(UriDb), &gorm.Config{})
	if err != nil {
		fmt.Println(err, "Connect DB failed err: %v ")
		return nil, err
	}

	return db, err
}

//func databasePingCheck(database *gorm.DB, timeout time.Duration) healthcheck.Check {
//	return func() error {
//		ctx, cancel := context.WithTimeout(context.Background(), timeout)
//		defer cancel()
//		if database == nil {
//			return fmt.Errorf("database is nil")
//		}
//		return database().DB()
//	}
//}
func main() {
	var database *sql.DB
	db, _ := connectToDatabase()
	database, _ = db.DB()
	health := healthcheck.NewHandler()
	health.AddReadinessCheck("check-database-connection", healthcheck.DatabasePingCheck(database, 1*time.Second))
	if err := http.ListenAndServe("0.0.0.0:8080", health); err != nil {
		return
	}
}
