package app

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Mubinabd/project_control/internal/pkg/config"
	"github.com/Mubinabd/project_control/internal/pkg/genproto/auth"
	"github.com/Mubinabd/project_control/internal/pkg/logger"
	st "github.com/Mubinabd/project_control/internal/pkg/postgres"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Js(cfg *config.Config) (*gin.Engine, error) {
	basepath := "/home/mubina/project_control"
	l := logger.NewLogger(basepath, cfg.LogPath)

	conn, err := st.New(cfg)
	if err != nil {
		l.ERROR.Printf("can't connect to db: %v", err)
		return nil, err
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/login", func(c *gin.Context) {
		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req := &auth.LoginReq{}
		err = json.Unmarshal(b, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `SELECT password, role FROM users WHERE username = $1`
		row := conn.DB.QueryRow(query, req.Username)
		var passwordHash, role string
		err = row.Scan(&passwordHash, &role)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not registered"})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		if role == "admin" {
			c.JSON(http.StatusAccepted, gin.H{"message": "Login successful", "redirect": "./front/group.html"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful", "redirect": "./front/private.html"})
			return
		}
	})
	
	return router, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
