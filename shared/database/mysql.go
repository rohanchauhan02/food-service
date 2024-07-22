package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/rohanchauhan02/food-service/shared/config"
)

// MysqlInterface is an interface represents methods that can be used to get mysql connection.
type MysqlInterface interface {
	OpenMysqlConn() (*gorm.DB, error)
}

// database is a struct that implements MysqlInterface.
type database struct {
	SharedConfig config.ImmutableConfigInterface
}

func (d *database) OpenMysqlConn() (*gorm.DB, error) {
	log.Println("Opening mysql connection...")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		d.SharedConfig.GetDatabaseUser(),
		d.SharedConfig.GetDatabasePassword(),
		d.SharedConfig.GetDatabaseHost(),
		d.SharedConfig.GetDatabasePort(),
		d.SharedConfig.GetDatabaseName(),
	)
	log.Printf("Start open mysql connection ...%s", connectionString)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Disable table name's pluralization globally
	// db.SingularTable(true)

	// Block global update for all tables to prevent accidentally update all records in the database without a where clause.
	db.BlockGlobalUpdate(true)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(60)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(30 * time.Second)
	// Enable Logger, show detailed log
	db.LogMode(true)

	return db, nil
}

// NewMysql returns a new instance of MysqlInterface.
func NewMysql(sharedConfig config.ImmutableConfigInterface) MysqlInterface {
	if sharedConfig == nil {
		panic("[PANIC] sharedConfig is nil, immutable config is required to create a new instance of MysqlInterface")
	}
	return &database{
		SharedConfig: sharedConfig,
	}
}
