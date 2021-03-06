package main

import (
	"log"
	"phone-number-normalizer/pgconx"
)

func main() {
	if pgconx.Init() != nil {
		log.Fatal("the database failed to connect")
	}

	_, err1 := pgconx.New("0401 05 06034")
	if err1 != nil {
		log.Fatal(err1)
	}

	phoneNumbers, err := pgconx.All("phone_numbers")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range phoneNumbers {
		v.Normalize()
		err := pgconx.Update(v.PhoneNumber, v.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}
