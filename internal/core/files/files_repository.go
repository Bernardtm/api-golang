package files

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type FilesRepository interface {
	GetAll() ([]FileResponse, error)
	GetByLink(link string) (FileResponse, error)
	Create(data FileRequest) (string, error)
	Update(data FileRequest) error
	Delete(link string) error
	Paginate(page, size int) ([]FileResponse, error)
}

type filesRepository struct {
	db *sql.DB
}

func NewFilesRepository(db *sql.DB) *filesRepository {
	return &filesRepository{
		db: db,
	}
}

func (r *filesRepository) GetAll() ([]FileResponse, error) {
	var models []FileResponse

	rows, err := r.db.Query("SELECT * FROM default_schema.files")

	if err != nil {
		return nil, errors.New("failed to retrive all files")
	}

	defer rows.Close()

	for rows.Next() {
		var model FileResponse

		if err := rows.Scan(&model.FileUUID, &model.Name, &model.Link, &model.Type, &model.Folder, &model.StatusUUID, &model.CreationDate, &model.ModificationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		models = append(models, model)
	}

	return models, nil
}

func (r *filesRepository) GetByLink(link string) (FileResponse, error) {
	var model FileResponse

	err := r.db.QueryRow("SELECT * FROM default_schema.files", link).Scan(&model.FileUUID, &model.Name, &model.Type, &model.Link, &model.Folder, &model.StatusUUID, &model.CreationDate, &model.ModificationDate)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return FileResponse{}, errors.New("file not found")
		}

		return FileResponse{}, errors.New("failed to retrive file")
	}

	return model, nil
}

func (r *filesRepository) Create(data FileRequest) (string, error) {
	var id string

	query := `INSERT INTO default_schema.files (
		file_name, file_link, file_folder, file_type, status_uuid
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING file_uuid`

	err := r.db.QueryRow(query, data.Name, data.Link, data.Folder, data.Type, data.StatusUUID).Scan(&id)

	if err != nil {
		log.Print(err)
		return "", errors.New("failed to create file")
	}

	return id, nil
}

func (r *filesRepository) Update(data FileRequest) error {
	_, err := r.db.Exec(`
		UPDATE default_schema.files
		SET file_name = $1,
		    file_link = $2,
		    file_folder = $3,
		    file_type = $4,
		    status_uuid = $5
		    modification_date = CURRENT_DATE
		    WHERE file_uuid = $1`, data.FileUUID, data.Name, data.Link, data.Folder, data.Type, data.StatusUUID)

	if err != nil {
		return errors.New("failed to update file")
	}

	return nil
}

func (r *filesRepository) Delete(link string) error {
	_, err := r.db.Exec(`DELETE FROM default_schema.files WHERE file_link = $1`, link)

	if err != nil {
		return errors.New("failed to delete file")
	}

	return nil
}

func (r *filesRepository) Paginate(page, size int) ([]FileResponse, error) {
	var models []FileResponse

	offset := (page - 1) * size

	rows, err := r.db.Query(`SELECT * FROM default_schema.files LIMIT $1 OFFSET $2`, size, offset)

	if err != nil {
		return nil, errors.New("failed to retrive all files")
	}

	defer rows.Close()

	for rows.Next() {
		var model FileResponse

		if err := rows.Scan(&model.FileUUID, &model.Name, &model.Type, &model.Link, &model.StatusUUID, &model.CreationDate, &model.ModificationDate); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		models = append(models, model)
	}

	return models, nil
}
