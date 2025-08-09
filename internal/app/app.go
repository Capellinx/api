package app

import (
	"api/internal/adapter/handler/http"
	"api/internal/adapter/repository"
	"api/internal/adapter/routes"
	"api/internal/usecases/company"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run() {
	db := InitDatabase()

	companyRepository := repository.NewCompanyRepositoryPostgres(db)
	companyUseCase := company.NewCreateCompanyUseCase(companyRepository)
	companyHandler := http.NewCompanyHandler(companyUseCase)

	app := gin.Default()
	routes.RegisterCompanyRoutes(app, companyHandler)

	log.Println("🚀 Server running on port 8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func InitDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	postgresURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return db
}
