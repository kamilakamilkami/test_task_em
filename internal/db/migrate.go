package db

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS subscriptions (
		id SERIAL PRIMARY KEY,
		service_name TEXT NOT NULL,
		price INTEGER NOT NULL,
		user_id UUID NOT NULL,
		start_date DATE NOT NULL,
		end_date DATE,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("Migrations applied successfully.")
	return nil
}
