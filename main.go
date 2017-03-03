package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

func main() {
	connectToDb()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "list.tmpl.html", nil)
	})

	router.Run(":" + port)
}

func connectToDb(){
	db, err := gorm.Open("postgres", "postgres://tkyhqchnzdkooo:6b0972b89563ae27e6329227da9ef134ea373175ff7ef9b237f0b59a4f170135@ec2-54-235-181-120.compute-1.amazonaws.com:5432/d6l87trf4piv5j")
	if err != nil{
		panic("failed to connect database")
	}
	defer db.Close()
}
