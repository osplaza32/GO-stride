package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github/osplaza32/gormgis/db"
	"github/osplaza32/gormgis/models"
	"net/http"
	"github.com/grindhold/gominatim"
)

func AllCiclovias(c *gin.Context) {
	fmt.Println("entra")
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	paradas:= models.Paradero{}
	db.First(&paradas)
	fmt.Println(paradas)
	var as []float64
	as = append(as,paradas.Location.Lat)
	as = append(as,paradas.Location.Lng)
	resp:= respo{}
	resp.Location = as
	resp.Id = paradas.ID
	resp.Nombre = paradas.Nombre
	resp.Code = paradas.Code
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": resp})

}
type respo struct {
	Id string
	Location []float64
	Nombre string
	Code string
}
func Nominatimctrl(c * gin.Context) {
	fmt.Print(c.Request.URL.Query().Get("Hola"))
	gominatim.SetServer("http://nominatim.cedeus.cl/nominatim/")
	//gominatim.SetServer("http://nominatim.openstreetmap.org/")

	qry := gominatim.SearchQuery{
		Q:"Talca",
		Addressdetails: true,
		Polygon:true,Bounded:true}
	result, _ := qry.Get() // Returns []gominatim.SearchResult
	fmt.Printf("%#v\n", result)


	rqry := gominatim.ReverseQuery{
		Lat: "-33.5058532",
		Lon: "-70.55716139999998",Zoom:10,AddressDetails:true}
	rresp, _ := rqry.Get()
	//fmt.Printf("Found %s\n", rresp.DisplayName)
	fmt.Printf("%#v\n", rresp)
	fmt.Print(c.GetPostForm("clave"))
	fmt.Print(c.Keys)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data":caca} )

}
