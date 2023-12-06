package util

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
}
