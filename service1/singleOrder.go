// pdf page no 397

package main

import (
	"fmt"
 
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	TotalBill int  `json:"totalBill"`
	SelectProduct string `json:"selectProduct"`
}

var db *gorm.DB
var err error

func initDB() {
	dsn := "host=localhost user=postgres password=pe+@l8900klh0 dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Order{})
}
 
func validateCustomer(fn http.HandlerFunc) http.HandlerFunc{

	return http.handlefunc(func(w http.ResponseWriter, r *http.Request)){
	customer := r.urlquery().get(customer)

    url := fmt.Println("")

	resp, _ := http.Get("http://localhost:8080/customers")

var customer
er:= json.newDecoder(resp.body).Decode(&this. customer)

if err != nil {
	fmt.prinln("No user present")
}
	}

	w.header().Set(content type), applicaton json

 
}


func validateProduct(fn http.HandlerFunc) http.HandlerFunc{

	return http.handlefunc(func(w http.ResponseWriter, r *http.Request)){
	product := r.urlquery().get(product)

    url := fmt.Println("")

	resp, _ := http.Get("http://localhost:8080/products")

var  product
er:= json.newDecoder(resp.body).Decode(&this.product)

if err != nil {
	fmt.prinln("No product present")
}
	}

	w.header().Set(content type), applicaton json

 
}



func createOrder(c *gin.Context) { 
  if !validateCustomer(){
	fmt.Println("Customer is not available")
	return
  } 
  
	if !validateProduct(){
		fmt.Println("Product is not available")
		return
	}  
	var order Order
		order.SelectProduct = c.Product
		order.TotalBill = c.Product.Amount
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := db.Create(&order); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func getOrder(c *gin.Context) {
	id := c.Param("id")
	var order Order
	if result := db.First(&order, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func main() {
	initDB()

	r := gin.Default()

	r.POST("/orders", createOrder)

	r.GET("/orders/:id", getOrder)

	r.Run(":8080")
}
