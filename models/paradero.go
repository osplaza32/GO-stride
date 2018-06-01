package models

import "github.com/nferruzzi/gormGIS"

type Paradero struct {
	ID       string           `gorm:"primary_key;uuid"`
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	Nombre   string
	Code     string
}
