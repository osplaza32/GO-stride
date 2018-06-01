package models

import "github.com/nferruzzi/gormGIS"

type PuntosCiclovias struct {
	ID        string `gorm:"primary_key;uuid"`
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	CivloviaID  uint     `gorm:"index"` // Foreign key (belongs to), tag `index` will create index for this column


}