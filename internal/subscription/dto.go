package subscription

type CreateSubscriptionDTO struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	UserID      string `json:"user_id" binding:"required,uuid"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

type FilterSumDTO struct {
	UserID      string `form:"user_id"`
	ServiceName string `form:"service_name"`
	FromDate    string `form:"from"`
	ToDate      string `form:"to"`
}
