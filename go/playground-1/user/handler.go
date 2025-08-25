package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func UserHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// 회원가입 핸들러
func (h *Handler) SignUpHandler(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"message": err.Error(),
		})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name is required",
		})
		return
	}
	if req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}

	response, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    response,
	})
}

// 사용자 조회
func (h *Handler) GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := h.service.GetUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to get user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data":    user,
	})
}

// 사용자 업데이트
func (h *Handler) UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	var req UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
	}

	_, err := h.service.UpdateUser(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
	}	

	 c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

// 사용자 삭제
func (h *Handler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	response, err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Failed to delete user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"data":    response,
	})
}

// 로그인
func (h *Handler) LoginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	response, err := h.service.Login(req)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"data":    response,
	})
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
	router.POST("/sign-up", h.SignUpHandler)
	router.POST("/login", h.LoginHandler)
	router.GET("/user/:id", h.GetUserHandler)
	router.PATCH("/user/:id", h.UpdateUserHandler)
	router.DELETE("/user/:id", h.DeleteUserHandler)
}


