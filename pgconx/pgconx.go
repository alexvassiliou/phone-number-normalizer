package pgconx

import (
	"fmt"
	"log"
	"os"
	"phone-number-normalizer/contact"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

// Init to initialise a postgresql connection
func Init() error {
	conxStr := pgConfig()
	var err error
	conn, err = pgx.Connect(conxStr)
	if err != nil {
		return err
	}
	return nil
}

// All lists all the phone numbers
func All(table string) ([]contact.Number, error) {
	var ret []contact.Number

	query := "SELECT * FROM " + table

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n contact.Number
		err := rows.Scan(&n.ID, &n.PhoneNumber)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}

	return ret, nil
}

// Update a database entry
func Update(value string, id int) error {
	queryStr := fmt.Sprintf("UPDATE phone_numbers SET number='%s' WHERE id=%d;", value, id)
	rows, err := conn.Query(queryStr)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func pgConfig() pgx.ConnConfig {
	var config pgx.ConnConfig
	config.Host = loadEnvVarialble("PG_HOST")
	config.User = loadEnvVarialble("PG_USER")
	config.Password = loadEnvVarialble("PG_PASSWORD")
	config.Database = loadEnvVarialble("PG_DBNAME")

	return config
}

func loadEnvVarialble(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
