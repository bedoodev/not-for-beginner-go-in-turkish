package main

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message") // 2025/07/15 13:56:28 This is a log message

	log.SetPrefix("INFO: ")
	log.Println("This is first info message.")  // INFO: 2025/07/15 13:56:47 This is first info message.
	log.Println("This is second info message.") // INFO: 2025/07/15 13:57:15 This is second info message.

	log.SetFlags(log.Ldate)
	log.Println("This is a log message with only date.") // INFO: 2025/07/15 This is a log message with only date.

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("This is a custom configured log message.") // INFO: 2025/07/15 14:02:49.671418 main.go:17: This is a custom configured log message.

	infoLogger.Println("This is an info logger.")   // INFO: 2025/07/15 14:05:46 main.go:22: This is an info logger.
	warnLogger.Println("This is a warning logger.") // WARNING: 2025/07/15 14:05:46 main.go:23: This is a warning logger.
	errorLogger.Println("This is an error logger.") // ERROR: 2025/07/15 14:05:46 main.go:24: This is an error logger.

	file, err := os.OpenFile("logging/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		errorLogger.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message.")

	infoLogger.SetOutput(file)
	infoLogger.Println("This is an info message.")

	warnLogger.SetOutput(file)
	warnLogger.Println("This is a warning message.")

	errorLogger.SetOutput(file)
	errorLogger.Println("This is an error message.")
}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)
