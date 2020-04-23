package pgconx

import (
	"log"
	"os"

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
func All(table string) ([]string, error) {
	var ret []string

	query := "SELECT * FROM " + table

	rows, err := conn.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var n string
		err := rows.Scan(nil, &n)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
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
