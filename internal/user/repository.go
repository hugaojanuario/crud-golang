package user

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(req CreateUserRequest, hashedPassword string) (*User, error) {
	query := `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, created_at, updated_at
    `
	user := &User{}
	err := r.db.QueryRow(query, req.Name, req.Email, hashedPassword).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return user, nil
}

func (r *Repository) findAll() ([]User, error){
	query := `
        SELECT id, name, email, created_at, updated_at
        FROM users
        ORDER BY id
    `
	rows, err := r.db.Query(query)
	if err != nil{
		return nil, fmt.Errorf("erro ao buscar usuarios: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next(){
		var u User
		if err := rows.Scan(&u.Name, &u.Email, &u.CreatedAt, *&u.UpdatedAt); err != nil{
			return nil, fmt.Errorf("Erro ao ler o usuario: %w", err)
		}
		users =append(users, u)
	}

	return users, nil
}