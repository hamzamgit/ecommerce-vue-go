package main

import (
	"backend/Database"
	router "backend/Router"
	"os"
	"time"
)

func SetTimeZone() {
	// Set the timezone to Australia
	var tzName = "Australia/Melbourne"
	os.Setenv("TZ", tzName)
	// Refresh the local timezone information
	time.LoadLocation(tzName)

}

func main() {
	SetTimeZone()
	Database.Connect()
	router.Serve()
}
