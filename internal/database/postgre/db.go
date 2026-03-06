package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewConnection(cfg Config) (*sql.DB, error) {
	strgConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", strgConn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir a conexão com o banco %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar com o banco %w", err)
	}

	return db, nil
}
