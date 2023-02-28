package main

import (
	"os"

	"todo-backend/database/sql"
	"todo-backend/src/handler"
	"todo-backend/src/repository"
	"todo-backend/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env.dev"); err != nil {
		panic(err)
	}

	// init DB
	db, err := sql.InitDB()
	if err != nil {
		panic(err)
	}

	// run migration
	if os.Getenv("DB_USERNAME") == "root" {
		migration := sql.Migration{Db: db.DB}
		migration.RunMigration()
	}

	// init repository
	repo := repository.Init(*db)

	// init usecase
	uc := usecase.Init(repo)

	// init handler
	router := gin.Default()
	rest := handler.Init(router, uc)
	rest.Run()
}
