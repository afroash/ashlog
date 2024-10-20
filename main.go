package ashlog

import (
	"log"
)

// LogInfo logs a message at the info level. usefull for debugging
func LogInfo(message string) {
	log.Printf("INFO: " + message)
}

// LogError logs a message at the error level. use for logging errors.
func LogError(message string) {
	log.Printf("ERROR: " + message)
}

// LogWarn logs a message at the warn level. use for logging warnings.
func LogWarn(message string) {
	log.Printf("WARN: " + message)
}

// LogFatal logs a message at the fatal level. use for logging fatal errors and exiting the program.
func LogFatal(message string) {
	log.Fatalf("FATAL: " + message)
}
