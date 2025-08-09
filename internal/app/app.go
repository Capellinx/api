package app

import (
	"api/internal/domain/usecases/company"
	"api/internal/infra/handler/http"
	postgres2 "api/internal/infra/persistence/postgres/company"
	"api/internal/infra/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run() {
	db := InitDatabase()

	companyRepository := postgres2.NewCompanyRepositoryPostgres(db)
	companyCreate := company.NewCreateCompanyUseCase(companyRepository)
	companyFetchAll := company.NewFetchAllCompanyUseCase(companyRepository)

	companyHandler := http.NewCompanyHandler(companyCreate, companyFetchAll)

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
