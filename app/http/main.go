package main

import (
	"os"

	"todo-backend/database/sql"
	"todo-backend/src/handler"
	"todo-backend/src/repository"
	"todo-backend/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo-backend/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

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

	docs.SwaggerInfo.Title = "Todo API"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	rest.Run()
}
