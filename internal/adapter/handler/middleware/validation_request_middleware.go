package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func ValidationRequestMiddleware[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload T

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(payload); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Dados da request inválidos",
				"details": err.Error(),
			})
			return
		}

		c.Set("payload", payload)
		c.Next()
	}
}
