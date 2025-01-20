package main

import (
	 
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Amount int `json: "amount"`
}
 
 

var db *gorm.DB
var err error

func initDB() {
	dsn := "host=localhost user=postgres password=pe+@l8900klh0 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Customer{})
	db.AutoMigrate(&Product{})
}



func createCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&customer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}


func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

 

func getCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer Customer
	if result := db.First(&customer, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}


func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product
	if result := db.First(&product, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
 

 

func main() {
	initDB()

	r := gin.Default()

	r.POST("/customers", createCustomer)
	 
	r.GET("/customers/:id", getCustomer)
	 
	r.POST("/products",createProduct)
	 
	r.GET("/products/:id", getProduct)

	r.Run(":8080")
}
