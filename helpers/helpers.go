package helpers

import "log"

// LogFatal logs errors
func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
