package auth

//структура пользователя
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Message struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
