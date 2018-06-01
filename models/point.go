package models

import (
	"github.com/nferruzzi/gormGIS"
	"time"
)

type Points struct {
	ID        string `gorm:"primary_key;uuid"`
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	CreatedAt time.Time
	Score float32
	UserID  string     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column

}