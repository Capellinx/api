package http

import (
	"api/internal/usecases/company"
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
	payload, exists := c.Get("payload")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payload não encontrado no contexto"})
		return
	}

	i, ok := payload.(company.CreateCompanyInputDTO)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tipo do payload inválido no contexto"})
		return
	}

	err := h.createUseCase.Execute(c.Request.Context(), i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
