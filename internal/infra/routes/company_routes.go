package routes

import (
	"api/internal/domain/usecases/company"
	"api/internal/infra/handler/http"
	"api/internal/infra/handler/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(r *gin.Engine, handler *http.CompanyHandler) {
	companies := r.Group("/company")
	{
		companies.POST(
			"",
			middlewares.ValidationRequestMiddleware[company.CreateCompanyInputDTO](),
			handler.Create,
		)
		companies.GET(
			"",
			handler.FetchAll,
		)
		companies.PATCH(
			"/:id",
			handler.Desactive,
		)
	}
}
