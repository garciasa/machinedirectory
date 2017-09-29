package database

import (
	"fmt"
	"machinedirectory/server/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database Type for store database structure
type Database struct{ Db *gorm.DB }

// New returns a database
func New(user, password, dbname string) (Database, error) {
	var d Database
	var err error
	connect := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
	d.Db, err = gorm.Open("mysql", connect)
	if err != nil {
		return d, err
	}

	return d, nil
}

// CreateStructure create scheme in database
func (d *Database) CreateStructure() {
	d.Db.AutoMigrate(&storage.Item{})
}

//Create insert elements in DB
func (d *Database) Create(item *storage.Item) error {
	err := d.Db.Create(item)
	if err != nil {
		return err.Error
	}

	return nil
}
