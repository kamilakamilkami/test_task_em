package subscription

type CreateSubscriptionDTO struct {
	ServiceName string `json:"service_name" binding:"required" example:"Netflix"`
	Price       int    `json:"price" binding:"required" example:"1500"`
	UserID      string `json:"user_id" binding:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate   string `json:"start_date" binding:"required" example:"2025-01-01"`
	EndDate     string `json:"end_date" example:"2025-06-01"`
}

type FilterSumDTO struct {
	UserID      string `form:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	ServiceName string `form:"service_name" example:"Netflix"`
	FromDate    string `form:"from" example:"2025-01-01"`
	ToDate      string `form:"to" example:"2025-12-31"`
}
