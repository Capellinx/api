package routes

import (
	"api/internal/adapter/handler/http"
	"api/internal/adapter/handler/middleware"
	"api/internal/usecases/company"

	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(r *gin.Engine, handler *http.CompanyHandler) {
	companies := r.Group("/company")
	{
		companies.POST(
			"",
			middleware.ValidationRequestMiddleware[company.CreateCompanyInputDTO](),
			handler.Create)
	}
}
