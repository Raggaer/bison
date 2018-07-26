package main

import (
	"log"
)

func main() {
	// Load config file
	config, err := loadConfig()
	log.Println(config, err)
	//	router := httprouter.New()

}
