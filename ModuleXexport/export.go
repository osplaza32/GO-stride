package export

import (
	"io/ioutil"
	"fmt"
	"github.com/paulmach/go.geojson"
	"github/osplaza32/gormgis/models"
	"github/osplaza32/gormgis/db"
	"github.com/nferruzzi/gormGIS"
	"github.com/artonge/go-gtfs"
)


func SeeCivlovias(s string)  {
	db,err := db.Conneccion()
	if err != nil {

		fmt.Println(err)

	}
	if !db.HasTable(&models.PuntosCiclovias{}) && !db.HasTable(&models.Ciclovias{}) {
		db.AutoMigrate(&models.PuntosCiclovias{})
		db.AutoMigrate(&models.Ciclovias{})

		raw, _ := ioutil.ReadFile(s)
	fc, err := geojson.UnmarshalFeatureCollection(raw)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _,v := range fc.Features{
		ciclo:= models.Ciclovias{}
		ciclo.Nombre =v.Properties["name"].(string)
		db.Save(&ciclo)

		for _,a:= range v.Geometry.LineString{
			m := models.PuntosCiclovias{
				Location: gormGIS.GeoPoint{
					Lat: a[1],
					Lng: a[0]}}
			m.CivloviaID = ciclo.Id
			db.Save(&m)
		}

	}
	}else
	{
		fmt.Println("tablas ya creadas")
	}

}
func SeeInterseccion(s string){
	db,err := db.Conneccion()
	if err != nil {

		fmt.Println(err)

	}
	if !db.HasTable(&models.Semafaro{}) {
		db.AutoMigrate(&models.Semafaro{})

		raw, _ := ioutil.ReadFile(s)
	fc, err := geojson.UnmarshalFeatureCollection(raw)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _,v := range  fc.Features{
		intercepcion := models.Semafaro{}
		arr,comuna,err :=intercepcion.Split(v.Properties["description"].(string))
		if err != nil{
			fmt.Println("todo mal")
			return
		}
		intercepcion.Intercepcion = arr
		intercepcion.Comuna = comuna
		intercepcion.Location =  gormGIS.GeoPoint{
			Lat: v.Geometry.Point[1],
			Lng: v.Geometry.Point[0]}
		db.Save(&intercepcion)

	}
	}else{
		fmt.Println("Tabla ya creada")
	}

}
func Gtfshelp(path string){
	db,err := db.Conneccion()
	if err != nil {

		fmt.Println(err)

	}
	if !db.HasTable(&models.Paradero{}) {
		db.AutoMigrate(&models.Paradero{})
		feed, err := gtfs.Load(path, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range feed.Stops {
			paradero := models.Paradero{Location: gormGIS.GeoPoint{
				Lat: v.Latitude,
				Lng: v.Longitude,
			}, Nombre: v.Name, Code: v.Code}
			fmt.Println(paradero)
			db.Save(&paradero)

		}
	}else{
		fmt.Println("Tabla creada")
	}




}





