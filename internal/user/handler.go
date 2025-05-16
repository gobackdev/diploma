package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	repo      Repository
	jwtSecret string
}

func NewHandler(repo Repository, jwtSecret string) *Handler {
	return &Handler{repo: repo, jwtSecret: jwtSecret}
}

type registerRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register создаёт нового пользователя и возвращает JWT в заголовке Authorization.
func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	exists, err := h.repo.IsUserExists(req.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hashing error"})
		return
	}

	if err := h.repo.CreateUser(req.Login, string(hash)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create error"})
		return
	}

	token, err := GenerateToken(req.Login, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{
		"message": "user registered",
	})
}

// Login проверяет учётные данные и возвращает JWT при успешной аутентификации.
func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.repo.FindByLogin(req.Login)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login or password"})
		return
	}

	token, err := GenerateToken(user.Login, h.jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"message": "user logged in"})
}
