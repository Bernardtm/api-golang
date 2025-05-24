package menus

import (
	"bernardtm/backend/internal/utils"
	"database/sql"
	"fmt"
)

// MenusRepository defines the interface for menu-related operations
type MenusRepository interface {
	GetAll() ([]MenusResponse, error)
	GetByID(id string) (MenusResponse, error)
	Create(entity MenusRequest) (string, error)
	Update(id string, entity MenusRequest) error
	Delete(id string) error
	Paginate(page, size int) ([]MenusResponse, error)
	GetMenusByUserID(userUUID string) ([]MenusResponse, error)
}

type menusRepository struct {
	db *sql.DB
}

// NewMenusRepository creates a new instance of MenusRepository
func NewMenusRepository(db *sql.DB) *menusRepository {
	return &menusRepository{db: db}
}

// GetAll retrieves all menus
func (r *menusRepository) GetAll() ([]MenusResponse, error) {
	var entities []MenusResponse
	query := `
		SELECT
			  menu_uuid
			, name
			, icon
			, url
			, order_index
			, creation_date
			, modification_date
			, status_uuid
		FROM
			default_schema.menus
		ORDER BY order_index asc
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all menus: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entity MenusResponse
		if err := rows.Scan(
			&entity.MenuUUID,
			&entity.Name,
			&entity.Icon,
			&entity.URL,
			&entity.OrderIndex,
			&entity.CreationDate,
			&entity.ModificationDate,
			&entity.StatusUUID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		convertedCreationDate, err := utils.ParseDateISO8601(entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return nil, err
		}

		entity.CreationDate = &convertedCreationDate
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

// GetByID retrieves a menu by its ID
func (r *menusRepository) GetByID(id string) (MenusResponse, error) {
	var entity MenusResponse
	query := `
		SELECT
			  menu_uuid
			, name
			, icon
			, url
			, order_index
			, creation_date
			, modification_date
			, status_uuid
		FROM
			default_schema.menus
		WHERE
			menu_uuid = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&entity.MenuUUID,
		&entity.Name,
		&entity.Icon,
		&entity.URL,
		&entity.OrderIndex,
		&entity.CreationDate,
		&entity.ModificationDate,
		&entity.StatusUUID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return MenusResponse{}, fmt.Errorf("menu not found with ID: %s", id)
		}
		return MenusResponse{}, fmt.Errorf("failed to get menu by ID: %w", err)
	}
	convertedCreationDate, err := utils.ParseDateISO8601(entity.CreationDate)

	if err != nil {
		fmt.Println("Erro ao parsear a data:", err)
		return MenusResponse{}, err
	}

	entity.CreationDate = &convertedCreationDate
	if entity.ModificationDate != nil {
		convertedModificationDate, err := utils.ParseDateISO8601(entity.ModificationDate)
		if err != nil {
			fmt.Println("Erro ao parsear a ModificationDate:", err)
			return MenusResponse{}, err
		}
		entity.ModificationDate = &convertedModificationDate
	} else {
		// Se ModificationDate for nil, continue sem fazer alterações
		entity.ModificationDate = nil
	}

	return entity, nil
}

// Create inserts a new menu and returns its UUID
func (r *menusRepository) Create(entity MenusRequest) (string, error) {
	var id string
	query := `
		INSERT INTO default_schema.menus (
			  name
			, icon
			, url
			, order_index
			, status_uuid
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING menu_uuid
	`
	err := r.db.QueryRow(query,
		entity.Name,
		entity.Icon,
		entity.Url,
		entity.OrderIndex,
		entity.StatusUUID,
	).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to create menu: %w", err)
	}
	return id, nil
}

// Update modifies an existing menu by its ID
func (r *menusRepository) Update(id string, entity MenusRequest) error {
	query := `
		UPDATE default_schema.menus
		SET
			name = $2,
			icon = $3,
			url = $4,
			order_index = $5,
			modification_date = CURRENT_DATE,
			status_uuid = $6
		WHERE menu_uuid = $1
	`
	_, err := r.db.Exec(query,
		id,
		entity.Name,
		entity.Icon,
		entity.Url,
		entity.OrderIndex,
		entity.StatusUUID,
	)
	if err != nil {
		return fmt.Errorf("failed to update menu: %w", err)
	}
	return nil
}

// Delete removes a menu by its ID
func (r *menusRepository) Delete(id string) error {
	query := `
		DELETE FROM default_schema.menus
		WHERE menu_uuid = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete menu: %w", err)
	}
	return nil
}

// Paginate retrieves menus with pagination
func (r *menusRepository) Paginate(page, size int) ([]MenusResponse, error) {
	var entities []MenusResponse
	offset := (page - 1) * size
	query := `
		SELECT
			menu_uuid
			, name
			, icon
			, url
			, order_index
			, creation_date
			, modification_date
			, status_uuid
		FROM
			default_schema.menus
		ORDER BY order_index asc
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.Query(query, size, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to paginate menus: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entity MenusResponse
		if err := rows.Scan(
			&entity.MenuUUID,
			&entity.Name,
			&entity.Icon,
			&entity.URL,
			&entity.OrderIndex,
			&entity.CreationDate,
			&entity.ModificationDate,
			&entity.StatusUUID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		convertedCreationDate, err := utils.ParseDateISO8601(entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return nil, err
		}

		entity.CreationDate = &convertedCreationDate
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

func (r *menusRepository) GetMenusByUserID(userUUID string) ([]MenusResponse, error) {
	var entities []MenusResponse
	query := `
		SELECT
			  m.menu_uuid
			, m.name
			, m.icon
			, m.url
			, m.order_index
			, m.creation_date
			, m.modification_date
			, m.status_uuid
		FROM
			default_schema.menus m
		WHERE
			m.menu_uuid IN (
				SELECT um.menu_uuid
				FROM default_schema.user_menus um
				WHERE um.user_uuid = $1
			)
		ORDER BY order_index asc
	`
	rows, err := r.db.Query(query, userUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user menus: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var entity MenusResponse
		if err := rows.Scan(
			&entity.MenuUUID,
			&entity.Name,
			&entity.Icon,
			&entity.URL,
			&entity.OrderIndex,
			&entity.CreationDate,
			&entity.ModificationDate,
			&entity.StatusUUID,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		convertedCreationDate, err := utils.ParseDateISO8601(entity.CreationDate)

		if err != nil {
			fmt.Println("Erro ao parsear a data:", err)
			return nil, err
		}

		entity.CreationDate = &convertedCreationDate
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
