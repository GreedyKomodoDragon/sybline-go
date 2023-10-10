package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"sync"
	"time"

	handler "github.com/GreedyKomodoDragon/sybline-go"
)

func main() {
	isProducer := flag.Bool("producer", true, "")
	running := flag.Bool("running", false, "if consumer is in loop")
	queue := flag.Bool("queue", false, "creates queue")
	amount := flag.Int("amount", 50, "number of consumers")
	flag.Parse()

	if *isProducer {
		producer(*queue, *amount)
		return
	}

	consumer(*running)
}

func producer(queue bool, amount int) {

	passwordManager := handler.NewUnsecurePasswordManager()
	passwordManager.SetPassword("sybline", "sybline")

	client, err := handler.NewTLSSyblineClient(
		[]string{"localhost:2221", "localhost:2222", "localhost:2223"},
		"cert/ca-cert.pem",
		"cert/cert.pem",
		"cert/key.pem",
		true,
		passwordManager,
		handler.Config{
			TimeoutSec:      5,
			TimeoutAttempts: 3,
		})

	if err != nil {
		log.Fatal("error cannot connect:", err)
	}

	ctx := context.Background()

	// this is the critical step that includes your headers
	if err := client.Login(ctx, "sybline"); err != nil {
		log.Fatal("error login:", err)
	}

	if queue {
		amount := uint32(100_000)
		queueName := "randomOne"
		queueNameTwo := "randomTwo"

		if err := client.CreateQueue(ctx, "route", queueName, amount, 3, true); err != nil {
			log.Fatal("error queue:", err)
		}

		if err := client.CreateQueue(ctx, "routeTwo", queueNameTwo, amount, 3, true); err != nil {
			log.Fatal("error queue:", err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	message := []handler.Message{
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "route",
			Data: []byte("data"),
		},
		{
			Rk:   "routeTwo",
			Data: []byte("data"),
		},
		{
			Rk:   "routeTwo",
			Data: []byte("data"),
		},
		{
			Rk:   "routeTwo",
			Data: []byte("data"),
		},
		{
			Rk:   "routeTwo",
			Data: []byte("data"),
		},
	}

	fmt.Println(len(message))

	for k := 0; k < amount; k++ {
		go func(i int, messages []handler.Message) {
			for {
				// startTime := time.Now()
				if err := client.SubmitBatchMessage(ctx, messages); err != nil {
					fmt.Println("error submit:", err, i)
				}
				// fmt.Println(time.Since(startTime))
				// time.Sleep(20 * time.Millisecond)
			}

		}(k, message)
	}
	wg.Wait()
}

func consumer(running bool) {
	flag.Parse()

	passwordManager := handler.NewUnsecurePasswordManager()
	passwordManager.SetPassword("sybline", "sybline")

	var wg sync.WaitGroup
	wg.Add(1)

	for h := 0; h < 20; h++ {
		client, err := handler.NewBasicSyblineClient([]string{"localhost:2221", "localhost:2222", "localhost:2223"}, passwordManager, handler.Config{
			TimeoutSec:      5,
			TimeoutAttempts: 3,
		})
		if err != nil {
			log.Fatal("error cannot connect:", err)
		}

		go func(client handler.SyblineClient, i int) {
			// this is the critical step that includes your headers
			if err := client.Login(context.Background(), "sybline"); err != nil {
				log.Fatal("error login:", err)
			}

			if running {
				queue := "randomOne"
				if i%2 == 0 {
					queue = "randomTwo"
				}
				consumer, err := client.Consumer(20, 100*time.Millisecond, queue)
				if err != nil {
					log.Fatal("error consumer:", err)
				}

				for {
					data := <-consumer.Messages
					dataTwo := <-consumer.Messages
					go func(data, dataTwo *handler.MessageData) {

						ids := [][]byte{data.Id, dataTwo.Id}
						if err := consumer.BatchAck(context.Background(), ids); err != nil {
							fmt.Println("batch ack err:", err, data.Id)
						}
					}(data, dataTwo)
				}
			}

			msgs, err := client.GetMessages(context.Background(), "randomOne", 10)
			if err != nil {
				log.Fatal("error get:", err)
			}

			for _, msg := range msgs {
				fmt.Println("id:", msg.Id)
			}

			time.Sleep(2 * time.Minute)

			// should be the same
			if err := client.Login(context.Background(), "sybline"); err != nil {
				log.Fatal("error login:", err)
			}

			msgs, err = client.GetMessages(context.Background(), "randomOne", 10)
			if err != nil {
				log.Fatal("error get:", err)
			}

			for _, msg := range msgs {
				if err := client.Ack(context.Background(), "randomOne", msg.Id); err != nil {
					fmt.Println("ack err:", err, msg.Id)
					continue
				}

				if err := client.Nack(context.Background(), "randomOne", msg.Id); err != nil {
					fmt.Println("nack err:", err, msg.Id)
					continue
				}
			}
		}(client, h)
	}

	wg.Wait()
}
