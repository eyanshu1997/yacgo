package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var enabled bool

func Println(str string) {
	_, file, no, _ := runtime.Caller(1)
	if enabled {
		log.Println(fmt.Sprintf(" %s %s"+str, file, no))
	}
}

func Printf(str string, args ...interface{}) {
	_, file, no, _ := runtime.Caller(1)
	if enabled {
		log.Printf(fmt.Sprintf(" %s %s"+str, file, no), args...)
	}
}

func init() {
	logEnabled, ok := os.LookupEnv("LOG_LEVEL")
	log.Printf("Log is enabled %t", ok)
	// LOG_LEVEL not set, let's default to debug
	if ok {
		if logEnabled == "true" || logEnabled == "TRUE" {
			enabled = true
		}
	}
}
