package common

import (
	"log"
)

func DealWithError(msg string, err error) {
	if err != nil {
		log.Printf("%s\n%v", msg, err)
	}
}

func DieWithError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s\n%v", msg, err)
	}
}
