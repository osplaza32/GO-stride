package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"net/http"
	"./controller"
	"./db"
	"./models"
	"github.com/gin-gonic/gin"
)



func main() {
	db,err := db.Conneccion()
	if err != nil {

		fmt.Println(err)

	}
	db.AutoMigrate(&models.User{})

	db.AutoMigrate(&models.Colletion{})

	db.AutoMigrate(&models.Points{})

	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("static/importdata/assert","./static")


	v0 := router.Group("/static/importdata")
	{

		v0.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}


	v1 := router.Group("/api/v1/stride")
	{
		v1.POST("/login", controller.Loginhadler)
		v1.POST("/user", controller.CreateUser)
		v1.GET("/users", controller.AllUser)
		v1.GET("/user/:id", controller.ThisUser)
		//v1.PUT("/user/:id", updateUser)
		//v1.DELETE("/user/:id", deleteUser)
		v1.GET("/point/:id", controller.Getpoint4user)
		v1.GET("/points", controller.Getallpoints)
		v1.POST("/point" , controller.CreatePoint)
		v1.GET("/point/:id/:idp", controller.GetespecificPoint)

	}

	v3 := router.Group("/api/cedeus/nominatim")
	{
		v3.GET("",controller.Nominatimctrl)
	}
	router.Run(":8000")




}



