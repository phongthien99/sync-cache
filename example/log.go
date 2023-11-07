package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// Create a log file for writing log messages.
	logFile, err := os.Create("app.log")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer logFile.Close()

	// Create a logger that writes to both the console and the log file.
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "APP: ", log.Ldate|log.Ltime)

	// Create a channel to receive log messages.
	logChannel := make(chan string)

	// Start a goroutine to continuously log messages.
	go func() {
		for {
			select {
			case logMessage := <-logChannel:
				logger.Println(logMessage)
			}
		}
	}()

	// Simulate some log messages being generated in real-time.
	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Log message %d", i)
		logChannel <- message
		time.Sleep(2 * time.Second) // Simulate some activity before the next log message.
	}

	// Give some time for the goroutine to log the last message before exiting.
	time.Sleep(2 * time.Second)
}
