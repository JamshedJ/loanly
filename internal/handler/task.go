package handler

import (
	"net/http"

	"github.com/JamshedJ/REST-api/internal/models"
	"github.com/gin-gonic/gin"
)



func (h *Handler) CreateTask(c *gin.Context) {
	userID := c.GetInt("user_id")

	var params models.TaskParams
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.services.CreateTask(c, userID, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	
}

func (h *Handler) GetTasks(c *gin.Context) {
	
}

func (h *Handler) UpdateTask(c *gin.Context) {
	
}

func (h *Handler) DeleteTask(c *gin.Context) {
	
}

func (h *Handler) MarkTask(c *gin.Context) {
	
}