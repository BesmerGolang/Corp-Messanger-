package auth

// отвечает за сохранение и поиск
import (
	"database/sql"
	"fmt"
	"log"
)

type Repository struct { // хранит базу данных
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
func (r *Repository) CreateUser(username string, passwordHash string) (User, error) {
	sqlQuery := `INSERT INTO users (username, password_hash) VALUES (?, ?)`
	result, err := r.db.Exec(sqlQuery, username, passwordHash)
	if err != nil {
		return User{}, fmt.Errorf("пользователь уже существует или ошибка БД: %v", err)
	}
	id, _ := result.LastInsertId()
	return User{
		ID:           int(id),
		Username:     username,
		PasswordHash: passwordHash,
	}, nil

}
func (r *Repository) GetUserByUsername(username string) (User, error) {
	var user User
	sqlQuery := `SELECT id, username, password_hash FROM users WHERE username = ?`
	err := r.db.QueryRow(sqlQuery, username).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (r *Repository) SaveMessage(userID int, content string) error {
	query := `INSERT INTO messages (user_id, content) VALUES (?, ?)`
	_, err := r.db.Exec(query, userID, content)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetLastMessages() ([]Message, error) {
	query := `SELECT id, user_id, content, created_at FROM (SELECT * FROM messages ORDER BY id DESC LIMIT 50) ORDER BY id ASC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message

	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.ID, &msg.UserID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			continue
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
