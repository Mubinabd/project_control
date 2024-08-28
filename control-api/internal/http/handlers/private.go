package handlers

import (
	"context"
	"strconv"

	pb "github.com/Mubinabd/project_control/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create private
// @Description Create a new private
// @Tags Private
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body pb.CreatePrivateReq true "private data"
// @Success 200 {string} string "message":"private created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/private/create [post]
func (h *Handler) CreatePrivate(c *gin.Context) {
	var req pb.CreatePrivateReq
	if err := c.ShouldBindJSON(&req); err != nil {

		h.Logger.ERROR.Println("Failed to bind request", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	_, err := h.Clients.Private.CreatePrivate(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create private:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Private created successfully"})
}

// @Summary Get Private
// @Description Get an private by ID
// @Tags Private
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Private ID"
// @Success 200 {object} pb.PrivateGet "Private data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Private not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/private/{id} [get]
func (h *Handler) GetPrivate(c *gin.Context) {
	req := pb.ById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Private.GetPrivate(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get private", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Private
// @Description Update an existing private by ID
// @Tags Private
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Private ID"
// @Param private body pb.UpdatePrivat true "Private update data"
// @Success 200 {string} string "message":"Private updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/private/update/{id} [put]
func (h *Handler) UpdatePrivate(c *gin.Context) {
	id := c.Param("id")
	var req pb.UpdatePrivat
	req.Id = id
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Clients.Private.UpdatePrivate(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "private updated successfully"})
}

// @Summary List Privaties
// @Description List privaties with filters
// @Tags Private
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.PrivateListRes "List of privaties"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/private/list [get]
func (h *Handler) ListPrivates(c *gin.Context) {
	var filter pb.PrivateListReq

	f := pb.Pagination{}
	filter.Pagination = &f

	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Pagination.Limit = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid limit", err)
			c.JSON(400, "Invalid limit value")
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Pagination.Offset = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid offset", err)
			c.JSON(400, "Invalid offset value")
			return
		}
	}

	resp, err := h.Clients.Private.ListPrivates(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list privaties", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Private
// @Description Delete an private by ID
// @Tags Private
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Private ID"
// @Success 200 {string} string "message":"Private deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/private/delete/{id} [delete]
func (h *Handler) DeletePrivate(c *gin.Context) {
	id := c.Param("id")

	req := &pb.DeletePrivat{Id: id}
	_, err := h.Clients.Private.DeletePrivate(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete private:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Private deleted successfully"})
}