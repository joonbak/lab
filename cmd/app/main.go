package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func createDir(currentDate string) {
	path := "./" + currentDate

	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal("Error creating directory: v%", err)
	}
}

func main() {
	currentDate := time.Now().Format("2006-01-02")
	createDir(currentDate)
	fmt.Println("Current Date: ", currentDate)
}
