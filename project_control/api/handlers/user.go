package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Mubinabd/project_control/api/docs"
	t "github.com/Mubinabd/project_control/api/token"
	auth "github.com/Mubinabd/project_control/pkg/genproto/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.UserRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user/profiles [get]
func (h *Handlers) GetProfile(c *gin.Context) {
	userID := getuserId(c)
	req := &auth.GetById{
		Id: userID,
	}

	profile, err := h.User.GetProfile(c, req)
	if err != nil {
		slog.Error("Error getting profile:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Retrieved profile")
	c.JSON(http.StatusOK, profile)
}

// EditProfile godoc
// @Summary Edit user profile
// @Description Update the profile of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param profile body auth.EditProfileReqBpdy true "Updated profile details"
// @Success 200 {object} string "Profile updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user/profiles [put]
func (h *Handlers) EditProfile(c *gin.Context) {
	userID := getuserId(c)

	var body auth.EditProfileReqBpdy
	if err := c.ShouldBindJSON(&body); err != nil {
		slog.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.UserRes{
		Id:          userID,
		Username:    body.Username,
		Email:       body.Email,
		FullName:    body.FullName,
		DateOfBirth: body.DateOfBirth,
	}

	input, err := json.Marshal(req)
	if err != nil {
		slog.Error("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = h.Producer.ProduceMessages("upd-user", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	slog.Info("Updated profile")
	c.JSON(http.StatusOK, "Profile updated successfully")
}

// ChangePassword godoc
// @Summary Change user password
// @Description Update the password of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param password body auth.ChangePasswordReqBody true "Updated password details"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user/passwords [put]
func (h *Handlers) ChangePassword(c *gin.Context) {
	userID := getuserId(c)

	var body auth.ChangePasswordReqBody
	if err := c.ShouldBindJSON(&body); err != nil {
		slog.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	password, err := t.HashPassword(body.NewPassword)
	if err != nil {
		slog.Error("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	body.NewPassword = password

	req := &auth.ChangePasswordReq{
		Id:              userID,
		CurrentPassword: body.CurrentPassword,
		NewPassword:     body.NewPassword,
	}

	input, err := json.Marshal(req)
	if err != nil {
		slog.Error("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	err = h.Producer.ProduceMessages("upd-pass", input)
	if err != nil {
		slog.Error("Error producing message:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Updated password")
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// GetSetting godoc
// @Summary Get user settings
// @Description Retrieve the settings of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} auth.Setting
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user/setting [get]
func (h *Handlers) GetSetting(c *gin.Context) {
	userID := getuserId(c)

	req := &auth.GetById{
		Id: userID,
	}

	setting, err := h.User.GetSetting(c, req)
	if err != nil {
		slog.Error("Error getting setting:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Retrieved setting")
	c.JSON(http.StatusOK, setting)
}

// EditSetting godoc
// @Summary Edit user settings
// @Description Update the settings of a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param setting body auth.Setting true "Updated setting details"
// @Success 200 {object} string "Setting updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user/setting [put]
func (h *Handlers) EditSetting(c *gin.Context) {
	userID := getuserId(c)

	var body auth.Setting
	if err := c.ShouldBindJSON(&body); err != nil {
		slog.Error("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.SettingReq{
		Id:           userID,
		PrivacyLevel: body.PrivacyLevel,
		Notification: body.Notification,
		Language:     body.Language,
		Theme:        body.Theme,
	}

	input, err := json.Marshal(req)
	if err != nil {
		slog.Error("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	err = h.Producer.ProduceMessages("upd-setting", input)
	if err != nil {
		slog.Error("Error producing message:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Updated setting")
	c.JSON(http.StatusOK, gin.H{"message": "Setting updated successfully"})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} string "User deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /v1/user [delete]
func (h *Handlers) DeleteUser(c *gin.Context) {
	userID := getuserId(c)

	req := &auth.GetById{
		Id: userID,
	}

	_, err := h.User.DeleteUser(c, req)
	if err != nil {
		slog.Error("Error deleting user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	slog.Info("Deleted user")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s deleted successfully", req.Id)})
}
