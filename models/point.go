package models

import (
	"github.com/nferruzzi/gormGIS"
	"time"
	"fmt"
	"../db"
)

type Points struct {
	ID        string `gorm:"primary_key;uuid"`
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	CreatedAt time.Time
	Score int32
	ColletionRefer  string     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column

}

func (p Points)GetUser()(u string){
	db,err := db.Conneccion()
	if err != nil {
		fmt.Println(err)
		}
	colle := Colletion{}
	db.Where("ID = ?",p.ColletionRefer).Find(&colle)
	return colle.UserID

}
