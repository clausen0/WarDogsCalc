package main

import (
	"fmt"
	"math"
)

type position struct {
	x float32
	y float32
}

func createExamplePositionPerson() position {
	return position{
		x: 67.701,
		y: 72.301,
	}
}

func createExamplePositionPing() position {
	return position{
		x: 67.850,
		y: 72.278,
	}
}

func calculatePingDistance(person, ping position) float32 {
	return float32(math.Sqrt(float64((ping.x-person.x)*(ping.x-person.x) + (ping.y-person.y)*(ping.y-person.y))))
}

func test() {
	person := createExamplePositionPerson()
	ping := createExamplePositionPing()

	distance := calculatePingDistance(person, ping)
	distanceMeters := distance * 1000 // Convert to meters
	fmt.Printf("Person: x = %.3f, y = %.3f\n", person.x, person.y)
	fmt.Printf("Ping: x = %.3f, y = %.3f\n", ping.x, ping.y)
	fmt.Printf("Distance: %f meters\n", distanceMeters)
}
