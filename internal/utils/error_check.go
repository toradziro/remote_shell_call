package utils

import (
	"log"
	"os"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Printf("Msg: %s Err: %v\n", msg, err)
		os.Exit(1)
	}
}
