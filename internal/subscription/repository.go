package subscription

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, s Subscription) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
         VALUES ($1, $2, $3, $4, $5)`,
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate)
	return err
}

func (r *Repository) GetAll(ctx context.Context) ([]Subscription, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions`)
	if err != nil {
		return nil, err
	}

	var subs []Subscription
	for rows.Next() {
		var s Subscription
		rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
		subs = append(subs, s)
	}
	return subs, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (Subscription, error) {
	var s Subscription
	err := r.db.QueryRowContext(ctx,
		`SELECT id, service_name, price, user_id, start_date, end_date
         FROM subscriptions WHERE id = $1`, id).
		Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	return s, err
}

func (r *Repository) Update(ctx context.Context, id int, s Subscription) error {
	_, err := r.db.ExecContext(ctx,
		`UPDATE subscriptions
         SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5
         WHERE id=$6`,
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, id)
	return err
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM subscriptions WHERE id=$1`, id)
	return err
}

func (r *Repository) GetSum(ctx context.Context,
	userID, service string, from, to time.Time) (int, error) {

	q := `
        SELECT SUM(price) 
        FROM subscriptions 
        WHERE start_date >= $1 AND (end_date <= $2 OR end_date IS NULL)
    `
	args := []any{from, to}

	if userID != "" {
		q += " AND user_id = $3"
		args = append(args, userID)
	}

	if service != "" {
		q += " AND service_name = $4"
		args = append(args, service)
	}

	var sum int
	err := r.db.QueryRowContext(ctx, q, args...).Scan(&sum)
	return sum, err
}
