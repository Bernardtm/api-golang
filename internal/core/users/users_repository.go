package users

import (
	"bernardtm/backend/internal/utils"
	"database/sql"
	"errors"
	"fmt"
)

// UserRepository define a interface para operações sobre usuários
type UserRepository interface {
	GetAll() ([]UserResponse, error)
	GetByID(id string) (UserResponse, error)
	Create(entity UserRequest) (string, error)
	Update(id string, entity UserRequest) error
	Delete(id string) error
	Paginate(page, size int) ([]UserResponse, error)
	GetByEmail(email string) (UserResponse, error)
}

type userRepository struct {
	db *sql.DB
}

// NewUserRepository cria uma nova instância de UserRepository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

// GetAll returns todos os usuários
func (r *userRepository) GetAll() ([]UserResponse, error) {
	var entities []UserResponse
	rows, err := r.db.Query(`
		SELECT user_uuid, username, email, tax_number, creation_date, modification_date, status_uuid
		FROM default_schema.users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entity UserResponse
		if err := rows.Scan(&entity.Id, &entity.Username, &entity.Email, &entity.TaxNumber, &entity.CreationDate, &entity.ModificationDate, &entity.StatusUUID); err != nil {
			return nil, err
		}
		convertedCreationDate, err := utils.ParseDateISO8601(&entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return nil, err
		}

		entity.CreationDate = convertedCreationDate
		if entity.ModificationDate != nil {
			convertedModificationDate, err := utils.ParseDateISO8601(entity.ModificationDate)
			if err != nil {
				fmt.Println("Erro ao parsear a ModificationDate:", err)
				return nil, err
			}
			entity.ModificationDate = &convertedModificationDate
		} else {
			// Se ModificationDate for nil, continue sem fazer alterações
			entity.ModificationDate = nil
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

// GetByID returns um usuário baseado no ID
func (r *userRepository) GetByID(id string) (UserResponse, error) {
	var entity UserResponse
	err := r.db.QueryRow(`
		SELECT user_uuid, username, email, tax_number, creation_date, modification_date, status_uuid, position, phone, profile_image_link
		FROM default_schema.users WHERE user_uuid = $1`, id).
		Scan(&entity.Id, &entity.Username, &entity.Email, &entity.TaxNumber, &entity.CreationDate, &entity.ModificationDate, &entity.StatusUUID, &entity.Position, &entity.Phone, &entity.ProfileImageLink)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.New("user not found")
		}
		convertedCreationDate, err := utils.ParseDateISO8601(&entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return UserResponse{}, err
		}

		entity.CreationDate = convertedCreationDate

		if entity.ModificationDate != nil {
			convertedModificationDate, err := utils.ParseDateISO8601(entity.ModificationDate)
			if err != nil {
				fmt.Println("Erro ao parsear a ModificationDate:", err)
				return UserResponse{}, err
			}
			entity.ModificationDate = &convertedModificationDate
		} else {
			// Se ModificationDate for nil, continue sem fazer alterações
			entity.ModificationDate = nil
		}

		return entity, err
	}
	return entity, nil
}

// Create creates a new usuário
func (r *userRepository) Create(entity UserRequest) (string, error) {
	var id string
	err := r.db.QueryRow(`
		INSERT INTO default_schema.users (
			username,
			email,
			"password",
			tax_number,
			status_uuid,
			position,
			phone
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING user_uuid`,
		entity.Username,
		entity.Email,
		entity.Password,
		entity.TaxNumber,
		entity.StatusUUID,
		entity.Position,
		entity.Phone,
	).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// Update altera os dados de um usuário
func (r *userRepository) Update(id string, entity UserRequest) error {
	_, err := r.db.Exec(`
		UPDATE default_schema.users
		SET username = $2,
			email = $3,
			password = $4,
			tax_number = $5,
			status_uuid = $6,
			modification_date = CURRENT_DATE
		WHERE user_uuid = $1`,
		id,
		entity.Username,
		entity.Email,
		entity.Password,
		entity.TaxNumber,
		entity.StatusUUID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete deleta um usuário by ID
func (r *userRepository) Delete(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM default_schema.users
		WHERE user_uuid = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

// Paginate returns uma lista de usuários paginada
func (r *userRepository) Paginate(page, size int) ([]UserResponse, error) {
	var entities []UserResponse
	offset := (page - 1) * size
	rows, err := r.db.Query(`
		SELECT user_uuid, username, email, tax_number, creation_date, modification_date, status_uuid
		FROM default_schema.users
		LIMIT $1 OFFSET $2`, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entity UserResponse
		if err := rows.Scan(&entity.Id, &entity.Username, &entity.Email, &entity.TaxNumber, &entity.CreationDate, &entity.ModificationDate, &entity.StatusUUID); err != nil {
			return nil, err
		}
		convertedCreationDate, err := utils.ParseDateISO8601(&entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return nil, err
		}

		entity.CreationDate = convertedCreationDate

		if entity.ModificationDate != nil {
			convertedModificationDate, err := utils.ParseDateISO8601(entity.ModificationDate)
			if err != nil {
				fmt.Println("Erro ao parsear a ModificationDate:", err)
				return nil, err
			}
			entity.ModificationDate = &convertedModificationDate
		} else {
			// Se ModificationDate for nil, continue sem fazer alterações
			entity.ModificationDate = nil
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

// GetByEmail returns an user by email
func (r *userRepository) GetByEmail(email string) (UserResponse, error) {
	var entity UserResponse
	err := r.db.QueryRow(`
		SELECT user_uuid, username, email, password, tax_number, creation_date, modification_date, status_uuid
		FROM default_schema.users WHERE email = $1`, email).
		Scan(&entity.Id, &entity.Username, &entity.Email, &entity.Password, &entity.TaxNumber, &entity.CreationDate, &entity.ModificationDate, &entity.StatusUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity, errors.New("user not found")
		}

		convertedCreationDate, err := utils.ParseDateISO8601(&entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return UserResponse{}, err
		}

		entity.CreationDate = convertedCreationDate
		if entity.ModificationDate != nil {
			convertedModificationDate, err := utils.ParseDateISO8601(entity.ModificationDate)
			if err != nil {
				fmt.Println("Erro ao parsear a ModificationDate:", err)
				return UserResponse{}, err
			}
			entity.ModificationDate = &convertedModificationDate
		} else {
			// Se ModificationDate for nil, continue sem fazer alterações
			entity.ModificationDate = nil
		}

		return entity, err
	}
	return entity, nil
}
