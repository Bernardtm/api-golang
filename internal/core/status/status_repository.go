package status

import (
	"database/sql"
	"fmt"
)

// StatusRepository defines the interface for status CRUD operations
type StatusRepository interface {
	GetAll() ([]StatusResponse, error)
	GetByID(id string) (StatusResponse, error)
	GetByName(name string) (StatusResponse, error)
	Create(entity StatusRequest) (string, error)
	Update(id string, entity StatusRequest) error
	Delete(id string) error
	Paginate(page int, size int) ([]StatusResponse, error)
}

type statusRepository struct {
	db *sql.DB
}

// NewStatusRepository creates a new instance of StatusRepository
func NewStatusRepository(db *sql.DB) *statusRepository {
	return &statusRepository{db: db}
}

// GetAll retrieves all status records
func (r *statusRepository) GetAll() ([]StatusResponse, error) {
	var statuses []StatusResponse
	rows, err := r.db.Query(`
		SELECT
			status_uuid, name, creation_date, modification_date
		FROM
			default_schema.status
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all statuses: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var status StatusResponse
		if err := rows.Scan(&status.StatusUUID, &status.Name, &status.CreationDate, &status.ModificationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}

// GetByID retrieves a status by its ID
func (r *statusRepository) GetByID(id string) (StatusResponse, error) {
	var status StatusResponse
	query := `
		SELECT
			status_uuid, name, creation_date, modification_date
		FROM
			default_schema.status
		WHERE status_uuid = $1
	`
	err := r.db.QueryRow(query, id).Scan(&status.StatusUUID, &status.Name, &status.CreationDate, &status.ModificationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return StatusResponse{}, fmt.Errorf("status not found: %w", err)
		}
		return StatusResponse{}, fmt.Errorf("failed to retrieve status by ID: %w", err)
	}
	return status, nil
}

func (r *statusRepository) GetByName(name string) (StatusResponse, error) {
	var status StatusResponse
	query := `
		SELECT
			status_uuid, name, creation_date, modification_date
		FROM
			default_schema.status
		WHERE name = $1
	`
	err := r.db.QueryRow(query, name).Scan(&status.StatusUUID, &status.Name, &status.CreationDate, &status.ModificationDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return StatusResponse{}, fmt.Errorf("status not found: %w", err)
		}
		return StatusResponse{}, fmt.Errorf("failed to retrieve status by ID: %w", err)
	}
	return status, nil
}

// Create inserts a new status and returns its UUID
func (r *statusRepository) Create(entity StatusRequest) (string, error) {
	var id string
	query := `
		INSERT INTO default_schema.status (
			name
		)
		VALUES ($1)
		RETURNING status_uuid
	`
	err := r.db.QueryRow(query, entity.Name).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to create status: %w", err)
	}
	return id, nil
}

// Update modifies an existing status by its ID
func (r *statusRepository) Update(id string, entity StatusRequest) error {
	_, err := r.db.Exec(`
		UPDATE default_schema.status
		SET
			name = $2,
			modification_date = CURRENT_DATE
		WHERE status_uuid = $1
	`, id, entity.Name)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}
	return nil
}

// Delete removes a status by its ID
func (r *statusRepository) Delete(id string) error {
	_, err := r.db.Exec(`
		DELETE FROM default_schema.status
		WHERE status_uuid = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete status: %w", err)
	}
	return nil
}

func (r *statusRepository) Paginate(page int, size int) ([]StatusResponse, error) {
	var entities []StatusResponse
	offset := (page - 1) * size
	query := `
		SELECT *
		FROM default_schema.status
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(query, size, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var entity StatusResponse
		if err := rows.Scan(&entity.StatusUUID, &entity.Name, &entity.CreationDate, &entity.ModificationDate, &entity.StatusUUID); err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
