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


}
func Getallpoints(c *gin.Context)  {
	var usuarios []models.Points
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Find(&usuarios)
	if len(usuarios) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Points found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": usuarios})






}
func CreatePoint(c *gin.Context)  {
	value, err := strconv.ParseFloat(c.PostForm("scorre"), 32)
	if err != nil {
		panic(fmt.Sprintf("%v", err))

	}
	score := float32(value)

	u := models.Points{Location: gormGIS.GeoPoint{
		Lat: 43.76857094631136,
		Lng: 11.292383687705296,
	},CreatedAt: time.Now(),
		Score: score, UserID: c.PostForm("idu")}
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.Create(&u)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created successfully!", "resourceId": u.ID})

}
func GetespecificPoint(c *gin.Context)  {

}