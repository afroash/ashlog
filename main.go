package ashlog

import (
	"fmt"
	"log"
)

// logMessage is a helper function that formats the message with an optional error
func logMessage(level, message string, err error) string {
	if err != nil {
		return fmt.Sprintf("%s: %s - %v", level, message, err)
	}
	return fmt.Sprintf("%s: %s", level, message)
}

// LogInfo logs a message at the info level. Useful for debugging.
func LogInfo(message string, err ...error) {
	log.Print(logMessage("INFO", message, getError(err)))
}

// LogError logs a message at the error level. Use for logging errors.
func LogError(message string, err ...error) {
	log.Print(logMessage("ERROR", message, getError(err)))
}

// LogWarn logs a message at the warn level. Use for logging warnings.
func LogWarn(message string, err ...error) {
	log.Print(logMessage("WARN", message, getError(err)))
}

// LogFatal logs a message at the fatal level. Use for logging fatal errors and exiting the program.
func LogFatal(message string, err ...error) {
	log.Fatal(logMessage("FATAL", message, getError(err)))
}

// getError is a helper function to safely get the error from the variadic parameter
func getError(err []error) error {
	if len(err) > 0 {
		return err[0]
	}
	return nil
}
