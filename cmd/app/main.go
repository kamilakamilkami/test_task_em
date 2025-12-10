// @title Subscription Service API
// @version 1.0
// @description API for managing user subscriptions.
// @BasePath /
// @schemes http

package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	migrations "test_task_em/internal/db"
	"test_task_em/internal/subscription"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "test_task_em/docs"
)

func waitForDB(dsn string) *sql.DB {
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil && db.Ping() == nil {
			log.Println("Database is ready!")
			return db
		}

		log.Println("Waiting for database to start...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Database did not start in time:", err)
	return nil
}

func main() {
	dsn := os.Getenv("DB_DSN")

	db := waitForDB(dsn)

	if err := migrations.RunMigrations(db); err != nil {
		log.Fatal("Migration error: ", err)
	}
	repo := subscription.NewRepository(db)
	service := subscription.NewService(repo)
	handler := subscription.NewHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
