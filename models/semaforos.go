package models

import (
	"github.com/nferruzzi/gormGIS"
	"strings"
	"errors"
)

type Semafaro struct {
	ID       string           `gorm:"primary_key;uuid"`
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	Intercepcion   string
    Comuna string

}
func (s Semafaro) Split(palabra string)(string,string,error)  {
	var arrstring string
	as := strings.Split(palabra,"-")
	if len(as) < 2 {
		return arrstring,"", errors.New("Minimum match not found")
	}else{
		comuna := strings.Split(as[len(as)-1],";")[1]
		for _,palabra := range as{
			arrstring += strings.Split(palabra,";")[0]

		}
		return arrstring,comuna,nil


	}
}