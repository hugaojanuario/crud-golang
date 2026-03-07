package user

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
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

func (r *Repository) FindAll() ([]User, error){
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
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil{
			return nil, fmt.Errorf("Erro ao ler o usuario: %w", err)
		}
		users =append(users, u)
	}

	return users, nil
}

func (r *Repository) FindByID(id int) (*User, error){
	query :=`SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1`

	user := &User{}
	err := r.db.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, nil //nao foi encontrado
	}
	if err != nil{
		return nil, fmt.Errorf("Erro ao buscar o usuario: %w", err)
	}

	return user, nil
}

func (r *Repository) Update(id int, req UpdateUserRequest)(*User, error){
	query := `UPDATE users
		SET name= $1, email = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, email, created_at, updated_at`

	user := &User{}
	err := r.db.QueryRow(query, req.Name, req.Email, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows{
		return nil, nil
	}
	if err != nil{
		return nil, fmt.Errorf("Erro ao atualizar o usuario: %w", err)
	}	

	return user, nil
}

func (r *Repository) Delete(id int) error {
    query := `DELETE FROM users WHERE id = $1`

    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("erro ao deletar usuário: %w", err)
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return sql.ErrNoRows // não encontrado
    }

    return nil
}
