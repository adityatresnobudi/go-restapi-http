package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeTable(db *sql.DB) error {
	q := `
		CREATE TABLE transactions (
  		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  		account_number VARCHAR(255) NOT NULL,
  		transaction_type VARCHAR(50) NOT NULL CHECK (transaction_type IN ('credit', 'debit')),
  		amount FLOAT NOT NULL, 
  		created_at TIMESTAMPTZ DEFAULT NOW(),
  		updated_at TIMESTAMPTZ DEFAULT NOW()
	);

	`
	if _, err := db.Exec(q); err != nil {
		log.Printf("initialize table error: %s\n", err.Error())
		return err
	}

	return nil
}
