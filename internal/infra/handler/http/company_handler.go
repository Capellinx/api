package http

import (
	"api/internal/domain/usecases/company"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyHandler struct {
	createUseCase *company.CreateCompanyUseCase
	findAll       *company.FetchAllCompanyUseCase
	desactive     *company.DesactiveCompanyUseCase
	findOne       *company.FindOneCompanyUseCase
}

func NewCompanyHandler(
	createUC *company.CreateCompanyUseCase,
	fa *company.FetchAllCompanyUseCase,
	dst *company.DesactiveCompanyUseCase,
	fo *company.FindOneCompanyUseCase) *CompanyHandler {
	return &CompanyHandler{
		createUseCase: createUC,
		findAll:       fa,
		desactive:     dst,
		findOne:       fo,
	}
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

func (h *CompanyHandler) FetchAll(c *gin.Context) {
	cp, err := h.findAll.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cp)
}

func (h *CompanyHandler) Desactive(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da empresa é obrigatório"})
		return
	}

	err := h.desactive.Execute(c.Request.Context(), company.DesactiveCompanyInputDTO{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *CompanyHandler) Find(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID da empresa é obrigatório"})
		return
	}

	co, err := h.findOne.Execute(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, co)
}
