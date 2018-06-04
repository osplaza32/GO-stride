package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github/osplaza32/gormgis/models"
	"github/osplaza32/gormgis/db"
	"github.com/nferruzzi/gormGIS"
	"time"
	"strconv"
)

func Getpoint4user(c *gin.Context)  {
	var cole []models.Colletion
	userID := c.Param("id")

	var as []string
	var points []models.Points
	var responde []ResponseGetallPoints
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Where("proyect = ? and user_id = ?","Stride",userID).Find(&cole)
	for _,v := range cole{
		as= append(as, v.ID)
	}

	db.Where("colletion_refer in (?)",as).Find(&points)
	for _,v:= range points{
		responde = append(responde,ResponseGetallPoints{Location:v.Location,CreatedAt:v.CreatedAt,Score:v.Score})

	}
	if len(points) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Points found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responde})


}
func Getallpoints(c *gin.Context)  {
	var cole []models.Colletion
	var as []string
	var points []models.Points
	var responde []ResponseGetallPoints
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Where("proyect = ?","Stride").Find(&cole)
	for _,v := range cole{
		as= append(as, v.ID)
	}

	db.Where("colletion_refer in (?)",as).Find(&points)
	for _,v:= range points{
		responde = append(responde,ResponseGetallPoints{Location:v.Location,CreatedAt:v.CreatedAt,Score:v.Score})

	}
	if len(points) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Points found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": responde})






}
type ResponseGetallPoints struct{
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
	CreatedAt time.Time
	Score int32

}
func CreatePoint(c *gin.Context)  {

	i, err := strconv.ParseInt(c.PostForm("scorre"), 10, 32)
	if err != nil {
		panic(err)
	}
	scoreee:= int32(i)
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	exist,couuid :=GetColletion(c.PostForm("idu"))
	u := models.Points{Location: gormGIS.GeoPoint{
		Lat: 43.76857094631136,
		Lng: 11.292383687705296,
	},CreatedAt: time.Now(),
		Score: scoreee }
	if !exist {
		u.ColletionRefer=couuid

		}else{
			cole :=models.Colletion{
				CreatedAt:time.Now(),
				UserID: c.PostForm("idu"),
				Type: "Punto"}
			db.Create(&cole)
			u.ColletionRefer=cole.ID


	}
	db.Create(&u)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created successfully!", "resourceId": u.ID})

}
func GetespecificPoint(c *gin.Context)  {

}
func GetColletion(s string )(bool,string) {
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	colle := models.Colletion{}
	db.Where(&models.Colletion{UserID: s, Proyect: "Stride"}).First(&colle)
	return colle.IsEmpty(),colle.ID




}