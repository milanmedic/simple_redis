package utils

import (
	"fmt"
	"log"
	"os"
)

/*
Helper function to exit the program. Log and exit code depend on why the function was called.
*/
func Exit(code int, msg string) {
	if code != 0 {
		log.Fatalf(fmt.Sprintf("Error: %s\n Code: %d", msg, code))
	} else {
		fmt.Println(msg)
	}

	os.Exit(code)
}
