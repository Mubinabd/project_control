package middlerware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	t "github.com/Mubinabd/project_control/api/token"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func NewAuth(enforce *casbin.Enforcer) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/swagger/") {
			ctx.Next()
			return
		}

		allow, err := CheckPermission(ctx.FullPath(), ctx.Request, enforce)

		if err != nil {
			valid, _ := err.(jwt.ValidationError)
			if valid.Errors == jwt.ValidationErrorExpired {
				RequireRefresh(ctx)
			} else {
				RequirePermission(ctx)
			}
		} else if !allow {
			RequirePermission(ctx)
		}
		ctx.Next()
	}

}

func CheckPermission(path string, r *http.Request, enforcer *casbin.Enforcer) (bool, error) {
	role, err := GetRole(r)
	if err != nil {
		log.Println("Error while getting role from token: ", err)
		return false, err
	}
	method := r.Method

	allowed, err := enforcer.Enforce(role, path, method)
	if err != nil {
		log.Println("Error while comparing role from csv list: ", err)
		return false, err
	}

	return allowed, nil
}

func InvalidToken(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Invalid token !!!",
	})
}

// RequirePermission handles responses for insufficient permissions
func RequirePermission(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": "Permission denied",
	})
}

// RequireRefresh handles responses for expired access tokens
func RequireRefresh(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": "Access token expired",
	})
}
