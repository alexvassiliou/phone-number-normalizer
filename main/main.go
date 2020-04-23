package main

import (
	"fmt"
	"log"
	"phone-number-normalizer/pgconx"
)

func main() {
	fmt.Println("connecting to postgres...")
	if pgconx.Init() != nil {
		log.Fatal("the database failed to connect")
	}

	phoneNumbers, err := pgconx.All("phone_numbers")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range phoneNumbers {
		fmt.Println(v)
	}
}
