package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"../models"
	"../db"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func AllUser(c *gin.Context) {
	var usuarios []models.User
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Select("id,first_name,last_name,email").Find(&usuarios)
	if len(usuarios) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": usuarios})





}
func ThisUser(c *gin.Context){
	var user models.User
	userID := c.Param("id")
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Select("id,first_name,last_name,email").Where("id = ?", userID).First(&user)

	if len(user.ID) == 0  {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Usuario found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})


}
func CreateUser(c *gin.Context){
	b,_:=mailexist(c.PostForm("email"))
	if b {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.PostForm("clave")), bcrypt.DefaultCost)

		u := models.User{FirstName: c.PostForm("nombre"),
			LastName: c.PostForm("apellido"), Email: c.PostForm("email"), Password: string(hashedPassword[:])}
		db, err := db.Conneccion()
		defer db.Close()
		if err != nil {
			fmt.Println(err)
		}
		db.Create(&u)

		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created successfully!", "resourceId": u.ID})
	}else{
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "ese email ya esta registrado"})
		}
	}
func Loginhadler(c *gin.Context){
	db,err :=db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		}
	b,u := mailexist(c.PostForm("email"))
	if !b{
		err = bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(c.PostForm("clave")))
		fmt.Println(err)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "clave incorecta"})

			}else{
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "true", "resourceId": u.ID})
			}
			}else{
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "ese email no esta registrado"})
		}
	}
func mailexist(o string) (bool,models.User) {
	db, err := db.Conneccion()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	var u models.User
	db.Where("email = ?", o).First(&u)
	return u.IsEmpty(),u

}
