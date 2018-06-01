package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"reflect"
	"strings"
	"github.com/pborman/uuid"
	"os"
	"github.com/subosito/gotenv"
)

func Conneccion() (*gorm.DB,error) {
	gotenv.Load()
	db, err := gorm.Open("postgres", "host="+os.Getenv("HOST")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DB")+" sslmode=disable password="+os.Getenv("PG_PASS"))
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", beforeCreate)

	db.LogMode(true)
	db.Exec("CREATE EXTENSION postgis")
	db.Exec("CREATE EXTENSION postgis_topology")

	return db,err
}
func beforeCreate(scope *gorm.Scope) {
	reflectValue := reflect.Indirect(reflect.ValueOf(scope.Value))
	if strings.Contains(string(reflectValue.Type().Field(0).Tag), "uuid") {
		uuid.SetClockSequence(-1)
		scope.SetColumn("id", uuid.NewUUID().String())
	}
}

