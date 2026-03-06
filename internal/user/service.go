package user

import (
    "database/sql"
    "errors"
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

type Service struct {
    repo *Repository
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) CreateUser(req CreateUserRequest) (*User, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("erro ao gerar hash da senha: %w", err)
    }

    user, err := s.repo.CreateUser(req, string(hashedPassword))
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *Service) GetAllUsers() ([]User, error) {
    return s.repo.FindAll()
}

func (s *Service) GetUserByID(id int) (*User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("usuário não encontrado")
    }

    return user, nil
}

func (s *Service) UpdateUser(id int, req UpdateUserRequest) (*User, error) {
    existing, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    if existing == nil {
        return nil, errors.New("usuário não encontrado")
    }

    return s.repo.Update(id, req)
}

func (s *Service) DeleteUser(id int) error {
    err := s.repo.Delete(id)
    if errors.Is(err, sql.ErrNoRows) {
        return errors.New("usuário não encontrado")
    }

    return err
}

//proximo handler