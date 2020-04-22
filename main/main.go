package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

func main() {
	conxStr := pgConfig()
	conn, err := pgx.Connect(conxStr)
	if err != nil {
		log.Fatal(err)
	}
	rows, err2 := conn.Query("SELECT * FROM phone_numbers")
	if err2 != nil {
		log.Fatal(err2)
	}
	for rows.Next() {
		var n string
		err := rows.Scan(nil, &n)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(n))
	}
}

func loadEnvVarialble(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func pgConfig() pgx.ConnConfig {
	var config pgx.ConnConfig
	config.Host = loadEnvVarialble("HOST")
	config.User = loadEnvVarialble("USER")
	config.Password = loadEnvVarialble("PASSWORD")
	config.Database = loadEnvVarialble("DBNAME")

	return config
}
