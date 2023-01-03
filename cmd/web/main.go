package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nikbat/bookings/cmd/pkg/config"
	"github.com/nikbat/bookings/cmd/pkg/handlers"
	"github.com/nikbat/bookings/cmd/pkg/helpers"
	"github.com/nikbat/bookings/cmd/pkg/render"
)

const numPool = 100
const portNumber = ":8080"

var s1 = "test"
var app config.AppConfig
var sessionManager *scs.SessionManager

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	//basic()
	firstPage()
}

func firstPage() {

	// update this when in production
	app.InProduction = false
	app.SameSite = http.SameSiteLaxMode

	//initialize teamplate cache
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Unable to create template cache")
	}
	//set teamplate cache to the global config
	app.TemplateCache = tc
	app.UseCache = false

	render.SetConfig(&app)

	repo := handlers.NewRepository(&app)
	handlers.SetRepository(repo)

	/*
		Old approach
	*/
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.ListenAndServe(portNumber, nil)

	//Inititlize a new session
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = app.SameSite
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	fmt.Println(fmt.Sprintf("Starting webserver at %s", portNumber))
	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server", err)
	}

}

func basic() {

	log.Println(s1)
	var s string
	s = "A"
	log.Println("original =>", s)
	changeString(s)
	log.Println("After changeString =>", s)
	changeStringWithPointer(&s)
	log.Println("After changeStringWithPointer =>", s)

	//Structs
	var test helpers.SomeType
	test.TypeName = "SomeName"
	log.Println(test.TypeName)

	intChan := make(chan int)
	defer close(intChan) // closing a channel after the function
	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)

	testjson := tojson()
	log.Println(testjson)

	parsejson(testjson)

}

func changeString(t string) {
	log.Println(t)
	t = "B"
	log.Println(t)
}

func changeStringWithPointer(t *string) {
	log.Println(t)
	*t = "B"
	log.Println(t)
}

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func tojson() string {
	var m1 Person
	m1.FirstName = "Nipu"
	m1.LastName = "Batra"
	m1.HairColor = "Black"
	m1.HasDog = false

	m2 := Person{"Saksh", "Gambhir", "Black", true}
	m3 := Person{"Samair", "Batr", "Black", true}

	var persons []Person
	persons = append(persons, m1)
	persons = append(persons, m2)
	persons = append(persons, m3)

	//log.Println(persons)

	newJson, err := json.MarshalIndent(persons, "", " ")

	if err != nil {
		log.Println(err)
	}
	return string(newJson)

}

func parsejson(testjson string) {
	var persons []Person
	err := json.Unmarshal([]byte(testjson), &persons)
	if err != nil {
		log.Println("Unmarshalling Json", err)
	}

	log.Println("unmarshallend: %v", persons)
}
