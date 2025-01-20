
package main

import (
	"fmt"
	"sync"
	"time" 
    "log"
    "github.com/rabbitmq/amqp091-go"
)

type BulkOrder struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	TotalBill int  `json:"totalBill"`
	SelectProduct []string `json:"selectProduct"`
}
 

func main() {
    conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }
    defer conn.Close()

    ch, err := conn.Channel() // open a channel for communication
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "order_queue", 
        true,           
        false,         
        false,        
        false,        
        nil,          
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }

	for i,value := range
    

	for i,_ := range SelectProduct{

		order := `{"id":"SelectProduct[i].product.Id","product":"SelectProduct[i].product.order"}`
		
		for j,_ := range order{
			order.TotalBill += order.SelectProduct.Amount
		}
	}

    err = ch.Publish(
        "",        
        q.Name,   
        false,     
        false,     
        amqp091.Publishing{
            ContentType: "application/json", 
            Body:        []byte(message),   
        },
    )
    if err != nil {
        log.Fatalf("Failed to publish a order: %s", err)
    }
    

  
}






 