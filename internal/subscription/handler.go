package subscription

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/subscriptions")
	{
		group.POST("/", h.Create)
		group.GET("/", h.List)
		group.GET("/:id", h.Get)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
		group.GET("/sum", h.Sum)
	}
}

// @Summary Create subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param data body CreateSubscriptionDTO true "Subscription data"
// @Success 201
// @Failure 400 {object} map[string]string
// @Router /subscriptions/ [post]
func (h *Handler) Create(c *gin.Context) {
	var dto CreateSubscriptionDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	start, _ := time.Parse("01-2006", dto.StartDate)
	var end *time.Time
	if dto.EndDate != "" {
		parsed, _ := time.Parse("01-2006", dto.EndDate)
		end = &parsed
	}

	s := Subscription{
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserID:      dto.UserID,
		StartDate:   start,
		EndDate:     end,
	}

	err := h.service.repo.Create(c, s)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary List all subscriptions
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} Subscription
// @Router /subscriptions/ [get]
func (h *Handler) List(c *gin.Context) {
	subs, err := h.service.repo.GetAll(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, subs)
}

// @Summary Get subscription by ID
// @Tags Subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} Subscription
// @Failure 404 {object} map[string]string
// @Router /subscriptions/{id} [get]
func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	sub, err := h.service.repo.GetByID(c, atoi(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, sub)
}

// @Summary Update subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param data body CreateSubscriptionDTO true "Updated subscription"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /subscriptions/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id := atoi(c.Param("id"))
	var dto CreateSubscriptionDTO
	if err := c.BindJSON(&dto); err != nil {
		return
	}

	start, _ := time.Parse("01-2006", dto.StartDate)
	var end *time.Time
	if dto.EndDate != "" {
		parsed, _ := time.Parse("01-2006", dto.EndDate)
		end = &parsed
	}

	s := Subscription{
		ServiceName: dto.ServiceName,
		Price:       dto.Price,
		UserID:      dto.UserID,
		StartDate:   start,
		EndDate:     end,
	}

	if err := h.service.repo.Update(c, id, s); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

// @Summary Delete subscription
// @Tags Subscriptions
// @Param id path int true "Subscription ID"
// @Success 200
// @Router /subscriptions/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := atoi(c.Param("id"))
	err := h.service.repo.Delete(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}

// @Summary Sum subscriptions by filter
// @Tags Subscriptions
// @Produce json
// @Param user_id query string false "User ID"
// @Param service_name query string false "Service name"
// @Param from query string false "From date"
// @Param to query string false "To date"
// @Success 200 {object} map[string]int
// @Router /subscriptions/sum [get]
func (h *Handler) Sum(c *gin.Context) {
	var q FilterSumDTO
	if err := c.BindQuery(&q); err != nil {
		return
	}

	from, _ := time.Parse("01-2006", q.FromDate)
	to, _ := time.Parse("01-2006", q.ToDate)

	sum, err := h.service.repo.GetSum(c, q.UserID, q.ServiceName, from, to)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"total": sum})
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
