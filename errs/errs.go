package errs

import (
	"log"
)

func CheckErr(str string, err error) {
	if err != nil {
		log.Print(str, err)
	}
}
