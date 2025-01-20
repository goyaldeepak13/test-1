package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)



func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
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

	orders, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to process order: %s", err)
	}

	fmt.Println("Waiting for order...")
	for j,order := range orders {
		var wg sync.WaitGroup 
			wg.Add(1)
			go Worker(j, jobs, results, &wg)
		
	}
}

// func Worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, job.ID)
// 		time.Sleep(time.Duration(job.Workload) * time.Millisecond) // Simulate work
// 		results <- fmt.Sprintf("Worker %d completed job %d", id, job.ID)
// 	}
// }

// func main() {
// 	const numWorkers = 3
// 	const numJobs = 10

// 	jobs := make(chan Job, numJobs)
// 	results := make(chan string, numJobs)

// var wg sync.WaitGroup

// for i := 1; i <= numWorkers; i++ {
// 	wg.Add(1)
// 	go Worker(i, jobs, results, &wg)
// }

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- Job{ID: j, Workload: 100 * j}
// 	}
// 	close(jobs)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println(result)
// 	}
// 	fmt.Printf("Total jobs processed: %d\n", numJobs)

// 	func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
// 		url := "amqp://" + MQHost + MQPort
// 		conn, err := amqp.Dial(url)
// 		if err != nil {
// 		  return nil, nil, err
// 		}
// 		channel, err := conn.Channel()
// 		if err != nil {
// 		  return nil, nil, err
// 		}

// 		if _, err := channel.QueueDeclare("", false, true, false, false,
// 	  nil); err != nil {
// 		  return nil, nil, err
// 		}
// 		return conn, channel, nil
// 	   }
