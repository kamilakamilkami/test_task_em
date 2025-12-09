package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"test_task_em/internal/subscription"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	repo := subscription.NewRepository(db)
	service := subscription.NewService(repo)
	handler := subscription.NewHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.Run(":8080")
}
