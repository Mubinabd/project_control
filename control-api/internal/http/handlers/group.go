package handlers

import (
	"context"
	"strconv"

	pb "github.com/Mubinabd/project_control/internal/pkg/genproto"

	"github.com/gin-gonic/gin"
)

// @Summary Create group
// @Description Create a new group
// @Tags Group
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body pb.CreateGroupReq true "group data"
// @Success 200 {string} string "message":"group created successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/group/create [post]
func (h *Handler) CreateGroup(c *gin.Context) {
	var req pb.CreateGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {

		h.Logger.ERROR.Println("Failed to bind request", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	_, err := h.Clients.Group.CreateGroup(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to create group:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Group created successfully"})
}

// @Summary Get Group
// @Description Get an group by ID
// @Tags Group
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} pb.GroupGet "Group data"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Group not found"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/group/{id} [get]
func (h *Handler) GetGroup(c *gin.Context) {
	req := pb.ById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Group.GetGroup(context.Background(), &req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to get group", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

// @Summary Update Group
// @Description Update an existing group by ID
// @Tags Group
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param group body pb.UpdateGr true "Group update data"
// @Success 200 {string} string "message":"Group updated successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/group/update/{id} [put]
func (h *Handler) UpdateGroup(c *gin.Context) {
	var req pb.UpdateGr
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Clients.Group.UpdateGroup(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "group updated successfully"})
}

// @Summary List Groups
// @Description List groups with filters
// @Tags Group
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.GroupListRes "List of groups"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/group/list [get]
func (h *Handler) ListGroups(c *gin.Context) {
	var filter pb.GroupListReq

	f := pb.Pagination{}
	filter.Pagination = &f

	if limit := c.Query("limit"); limit != "" {
		if value, err := strconv.Atoi(limit); err == nil {
			filter.Pagination.Limit = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid limit", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if value, err := strconv.Atoi(offset); err == nil {
			filter.Pagination.Offset = int32(value)
		} else {
			h.Logger.ERROR.Println("Invalid offset", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.Clients.Group.ListGroups(context.Background(), &filter)
	if err != nil {
		h.Logger.ERROR.Println("Failed to list groups", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

// @Summary Delete Group
// @Description Delete an group by ID
// @Tags Group
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {string} string "message":"Group deleted successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /v1/group/delete/{id} [delete]
func (h *Handler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")

	req := &pb.DeleteGr{Id: id}
	_, err := h.Clients.Group.DeleteGroup(context.Background(), req)
	if err != nil {
		h.Logger.ERROR.Println("Failed to delete group:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Group deleted successfully"})
}