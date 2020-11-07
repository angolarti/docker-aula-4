package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/wesleywillians/go-rabbitmq/queue"
	"html/template"
	"log"
	"net/http"
)

type Order struct {
	Coupon string
	CccNumber string
}

type Result struct {
	Status string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "<h1>Ola</h1>")
	tmpt := template.Must(template.ParseFiles("templates/home.html"))
	tmpt.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {

	coupon := r.PostFormValue("coupon")
	ccNumber := r.PostFormValue("cc-number")

	order := Order{
		Coupon: coupon,
		CccNumber: ccNumber,
	}

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error parsing to json")
	}
	// result := makeHttpCall("http://127.0.0.1:9091", r.FormValue("coupon"), r.FormValue("cc-number"))
	rabbitMQ := queue.NewRabbitMQ()
	ch := rabbitMQ.Connect()
	defer ch.Close()

	err = rabbitMQ.Notify(string(jsonOrder), "application/json", "orders_ex", "")
	if err != nil {
		log.Fatal("Error send message to the queue")
	}

	tmpt := template.Must(template.ParseFiles("templates/process.html"))
	tmpt.Execute(w, "")
}
