package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	t "github.com/Mubinabd/project_control/api/token"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := (ctx.Request.URL.Path)

		if strings.Contains(url, "swagger") || (url == "/auth/login") || (url == "/auth/register") {
			ctx.Next()
			return
		}
		ctx.Next()
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// if !strings.HasPrefix(authHeader, "Bearer ") {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		// 	c.Abort()
		// 	return
		// }

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		valid, err := t.ValidateToken(tokenString)
		if err != nil || !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		claims, err := t.ExtractClaim(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func GetUserId(r *http.Request) (string, error) {
	jwtToken := r.Header.Get("Authorization")

	if jwtToken == "" || strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}

	// if !strings.HasPrefix(jwtToken, "Bearer ") {
	// 	return "unauthorized", errors.New("invalid authorization header format")
	// }

	// tokenString := strings.TrimPrefix(jwtToken, "Bearer ")

	claims, err := t.ExtractClaim(jwtToken)
	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "unauthorized", errors.New("user_id claim not found")
	}
	return userID, nil
}

func GetRole(r *http.Request) (string, error) {
	jwtToken := r.Header.Get("Authorization")

	if jwtToken == "" || strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}

	claims, err := t.ExtractClaim(jwtToken)
	if err != nil {
		log.Println("Error while extracting claims: ", err)
		return "unauthorized", err
	}

	userID, ok := claims["role"].(string)
	if !ok {
		return "unauthorized", errors.New("role claim not found")
	}
	return userID, nil
}

func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Invalid token !!!",
	})
}

func RequirePermission(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Permission denied",
	})
}

func RequireRefresh(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Access token expired",
	})
}
