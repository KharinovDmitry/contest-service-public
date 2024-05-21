package middleware

import (
	"contest/internal/domain/enum"
	"contest/internal/domain/service/logger"
	"contest/lib/auth"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(logger logger.Logger, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Auth middleware get request")

		bearerHeader := c.Request.Header.Get("Authorization")
		if bearerHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(bearerHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims, err := auth.GetPayloadAndValidate(token, jwtSecret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(c.Request.Context(), enum.ContextKeyUsername, claims.Username)
		ctx = context.WithValue(ctx, enum.ContextKeyRole, claims.Role)
		ctx = context.WithValue(ctx, enum.ContextKeyID, int(claims.Id))

		c.Set(enum.ContextKeyUsername, claims.Username)
		c.Set(enum.ContextKeyRole, claims.Role)
		c.Set(enum.ContextKeyID, int(claims.Id))

		logger.Info(fmt.Sprintf("Claims: username:%s, role:%s, id:%f, id as int: %d", claims.Username, claims.Role, claims.Id, int(claims.Id)))

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
