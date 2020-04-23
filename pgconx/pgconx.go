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

// New insert a new number into the database
func New(number string) (int, error) {
	queryStr := fmt.Sprintf("INSERT INTO phone_numbers (number) VALUES ('%s')", number)

	// insert the new number
	rows, err := conn.Query(queryStr)
	if err != nil {
		return -1, err
	}
	rows.Close()

	// retrieve the new id
	ret, err2 := getID(number)
	if err2 != nil {
		return -1, nil
	}
	defer rows.Close()

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

func getID(number string) (int, error) {
	var ret int

	getID := fmt.Sprintf("SELECT id FROM phone_numbers WHERE number='%s'", number)

	rows, err := conn.Query(getID)
	if err != nil {
		return -1, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ret, nil)
		if err != nil {
			return -1, nil
		}
	}
	return ret, nil
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
