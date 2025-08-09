package http

import (
	"api/internal/usecases/company"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyHandler struct {
	createUseCase *company.CreateCompanyUseCase
}

func NewCompanyHandler(createUC *company.CreateCompanyUseCase) *CompanyHandler {
	return &CompanyHandler{createUseCase: createUC}
}

func (h *CompanyHandler) Create(c *gin.Context) {
	var payload company.CreateCompanyInputDTO

	fmt.Println(payload, c.Request.Body)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "corpo da requisição inválido"})
		return
	}

	err := h.createUseCase.Execute(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
