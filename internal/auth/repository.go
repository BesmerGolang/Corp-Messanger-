package auth

// отвечает за сохранение и поиск
import (
	"database/sql"
	"fmt"
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
