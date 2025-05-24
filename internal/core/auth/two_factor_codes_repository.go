package auth

import (
	"database/sql"
	"fmt"
)

// TwoFactorCodesRepository defines the interface for two-factor codes operations
type TwoFactorCodesRepository interface {
	GetByID(id string) (TwoFactorCodesResponse, error)
	Create(userid string, code string, expiryTime int) (string, error)
	Update(id string) error
}

type twoFactorCodesRepository struct {
	db *sql.DB
}

// NewTwoFactorCodesRepository creates a new instance of TwoFactorCodesRepository
func NewTwoFactorCodesRepository(db *sql.DB) *twoFactorCodesRepository {
	return &twoFactorCodesRepository{db: db}
}

// GetByID retrieves a two-factor code by its ID
func (r *twoFactorCodesRepository) GetByID(id string) (TwoFactorCodesResponse, error) {
	var entity TwoFactorCodesResponse
	err := r.db.QueryRow(`
		SELECT
			two_factor_code_uuid,
			user_uuid,
			code,
			expiration_date,
			status_uuid
		FROM default_schema.two_factor_codes
		WHERE two_factor_code_uuid = $1`, id).
		Scan(&entity.TwoFactorCodeUUID, &entity.Id, &entity.Code, &entity.ExpirationDate, &entity.StatusUUID)

	if err != nil {
		if err == sql.ErrNoRows {
			return TwoFactorCodesResponse{}, fmt.Errorf("two-factor code not found")
		}
		return TwoFactorCodesResponse{}, err
	}
	return entity, nil
}

// Create inserts a new two-factor code and returns its ID
func (r *twoFactorCodesRepository) Create(userid string, code string, minutesToExpiry int) (string, error) {
	var twoFactorCodeID string

	err := r.db.QueryRow(`
		INSERT INTO default_schema.two_factor_codes (
			user_uuid,
			code,
			expiration_date,
			status_uuid
		)
		VALUES ($1, $2, CURRENT_TIMESTAMP + $3::interval,
			(SELECT status_uuid FROM default_schema.status WHERE name = 'Actived'))
		RETURNING two_factor_code_uuid`,
		userid,
		code,
		fmt.Sprintf("%d minutes", minutesToExpiry)).
		Scan(&twoFactorCodeID)

	if err != nil {
		return "", err
	}
	return twoFactorCodeID, nil
}

// Update sets the status of a two-factor code to 'Inactived'
func (r *twoFactorCodesRepository) Update(id string) error {
	_, err := r.db.Exec(`
		UPDATE default_schema.two_factor_codes
		SET
			status_uuid = (SELECT status_uuid FROM default_schema.status WHERE name = 'Inactived'),
			modification_date = CURRENT_TIMESTAMP
		WHERE two_factor_code_uuid = $1`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
