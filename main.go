package ashlog

import (
	"fmt"
	"log"
)

// logMessage is a helper function that formats the message with an optional error
func logMessage(level, message string, err interface{}) string {
	if err != nil {
		return fmt.Sprintf("%s: %s - %v", level, message, err)
	}
	return fmt.Sprintf("%s: %s", level, message)
}

// LogInfo logs a message at the info level. Useful for debugging.
func LogInfo(message string, args ...interface{}) {
	var err interface{}
	if len(args) > 0 {
		err = args[0]
	}
	log.Print(logMessage("INFO", message, err))
}

// LogError logs a message at the error level. Use for logging errors.
func LogError(message string, args ...interface{}) {
	var err interface{}
	if len(args) > 0 {
		err = args[0]
	}
	log.Print(logMessage("ERROR", message, err))
}

// LogWarn logs a message at the warn level. Use for logging warnings.
func LogWarn(message string, args ...interface{}) {
	var err interface{}
	if len(args) > 0 {
		err = args[0]
	}
	log.Print(logMessage("WARN", message, err))
}

// LogFatal logs a message at the fatal level. Use for logging fatal errors and exiting the program.
func LogFatal(message string, args ...interface{}) {
	var err interface{}
	if len(args) > 0 {
		err = args[0]
	}
	log.Fatal(logMessage("FATAL", message, err))
}
