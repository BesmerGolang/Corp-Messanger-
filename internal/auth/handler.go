package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-key-for-my-chat-app-12345")

type Handler struct {
	repo *Repository // экземпляр структуры Repository для которой я писал метод CreateUser в repositiry
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Register(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Данные не соответствую формату JSON"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка хэширования пароля"})
		return
	}
	user, err := h.repo.CreateUser(req.Username, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Не удалось создать пользователя"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно создан", "user": user})
}
func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
func (h *Handler) Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Данные не соответствую формату JSON"})
		return
	}
	user, err := h.repo.GetUserByUsername(req.Username)
	if err != nil || !checkPassword(user.PasswordHash, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})
		return
	}
	token, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Успешный вход", "user_id": user.ID, "token": token})
}
