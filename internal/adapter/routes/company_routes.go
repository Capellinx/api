package routes

import (
	"api/internal/adapter/handler/http"

	"github.com/gin-gonic/gin"
)

func RegisterCompanyRoutes(r *gin.Engine, handler *http.CompanyHandler) {
	companies := r.Group("/company")
	{
		companies.POST("", handler.Create)
	}
}
