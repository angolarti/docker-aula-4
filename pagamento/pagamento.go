package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"github.com/wesleywillians/go-rabbitmq/queue"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Order struct {
	ID uuid.UUID
	Coupon string
	CcNumber string
}

type Result struct {
	Status string
}

func NewOrder() Order {
	uid, _ := uuid.NewV4()
	return Order{ID: uid}
}

const (
	InvalidCoupon = "invalid"
	ValidCoupon = "valid"
	ConnectionError = "connection error"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	messageChannel := make(chan amqp.Delivery)

	rabbitmq := queue.NewRabbitMQ()
	ch := rabbitmq.Connect()
	defer ch.Close()
	rabbitmq.Consume(messageChannel)

	for msg := range messageChannel {
		process(msg)
	}
}

func process(msg amqp.Delivery) {

	order := NewOrder()
	json.Unmarshal(msg.Body, &order)

	resultCoupon := makeHttpCall("http://127.0.0.1:9092", order.Coupon)

	switch resultCoupon.Status {
	case InvalidCoupon:
		log.Println("Order: ",order.ID, ": invalid coupon")
	case ConnectionError:
		msg.Reject(false)
		log.Println("Order: ",order.ID, ": could not process!")
	case ValidCoupon:
		log.Println("Order: ",order.ID, ": processed")
	}
}

func makeHttpCall(urlMicroservice string, coupon string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)

	res, err := http.PostForm(urlMicroservice, values) // http.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: ConnectionError}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error processing result")
	}

	result := Result{}
	json.Unmarshal(data, &result)

	return result
}


