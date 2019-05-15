package main

import (
	"fmt"
	"time"
)

// Exercise: run this code and try fixing the errors!

type Location struct {
	Lat float64
	Lng float64
}

type Festival struct {
	Location
	Name string
	Date time.Time
}

func main() {
	festivalDate := time.Date(1969, 8, 15, 12, 0, 0, 0, time.UTC)

	// Create and initialize a variable of type festival with a location
	woodstock := Festival{
		name:     "Woodstock 1969",
		location: location{lat: 41.371731, lng: -73.407448},
		date:     festivalDate,
	}
	fmt.Println("Festival name: ", festival.name)
	fmt.Printf("Festival location: (%f, %f)\n", festival.location.lat, festival.location.lng)
	fmt.Println("Festival date: ", festival.date)

	// Declare a variable using an anonymous struct
	person := struct {
		name string
		age  int
	}{
		name: "Janis",
		age:  25,
	}
	fmt.Println(person)
	fmt.Println(person.surname)
}
