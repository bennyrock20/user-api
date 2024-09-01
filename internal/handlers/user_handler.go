package handler

import (
	"net/http"
	"strconv"
	"taxi-service/internal/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *user.UserService
}

func NewUserHandler(service *user.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var u struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error input": err.Error()})
		return
	}
	if err := h.UserService.CreateUser(u.Name, u.Email, u.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error creating user": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.UserService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.UserService.GetUserByID(uint(id))
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UserService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "user updated"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.UserService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
}
