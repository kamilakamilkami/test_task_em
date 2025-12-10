package subscription

import "time"

// Subscription represents user subscription
// @Description Subscription model
type Subscription struct {
	ID          int        `json:"id" example:"1"`
	ServiceName string     `json:"service_name" example:"Netflix"`
	Price       int        `json:"price" example:"1500"`
	UserID      string     `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate   time.Time  `json:"start_date" example:"2025-01-01T00:00:00Z"`
	EndDate     *time.Time `json:"end_date,omitempty" example:"2025-06-01T00:00:00Z"`
}
