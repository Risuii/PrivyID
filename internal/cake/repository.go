package cake

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"privyID/helpers/exception"
	"privyID/models"
)

type (
	CakeRepository interface {
		Create(ctx context.Context, params models.CheeseCake) (int64, error)
		FindByID(ctx context.Context, id int64) (models.CheeseCake, error)
		FindAll() ([]models.CheeseCake, error)
		Update(ctx context.Context, id int64, cake models.CheeseCake) error
		Delete(ctx context.Context, id int64) error
	}

	cakeRepositoryImpl struct {
		db        *sql.DB
		tableName string
	}
)

func NewCakeRepository(db *sql.DB, tableName string) CakeRepository {
	return &cakeRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (repo *cakeRepositoryImpl) Create(ctx context.Context, params models.CheeseCake) (int64, error) {
	query := fmt.Sprintf(`INSERT INTO %s (title, description, rating, image, created_at) VALUES(?,?,?,?,?)`, repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		params.Title,
		params.Description,
		params.Rating,
		params.Image,
		params.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return 0, exception.ErrInternalServer
	}

	ID, _ := result.LastInsertId()

	return ID, nil
}

func (repo *cakeRepositoryImpl) FindByID(ctx context.Context, id int64) (models.CheeseCake, error) {
	var cake models.CheeseCake

	query := fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s WHERE id = ?`, repo.tableName)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return cake, exception.ErrInternalServer
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	err = row.Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdateAt,
	)
	if err != nil {
		log.Println(err)
		return cake, exception.ErrNotFound
	}

	return cake, nil
}

func (repo *cakeRepositoryImpl) FindAll() ([]models.CheeseCake, error) {
	var cakes []models.CheeseCake

	rows, err := repo.db.Query(fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s ORDER BY rating DESC, title ASC`, repo.tableName))
	if err != nil {
		log.Println(err)
		return cakes, exception.ErrInternalServer
	}

	defer rows.Close()

	for rows.Next() {
		var c models.CheeseCake
		if err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.Description,
			&c.Rating,
			&c.Image,
			&c.CreatedAt,
			&c.UpdateAt,
		); err != nil {
			log.Println(err)
			return cakes, exception.ErrInternalServer
		}
		cakes = append(cakes, c)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return cakes, exception.ErrInternalServer
	}

	return cakes, nil
}

func (repo *cakeRepositoryImpl) Update(ctx context.Context, id int64, cake models.CheeseCake) error {
	query := fmt.Sprintf(`UPDATE %s SET title = ?, description = ?, rating = ?, image = ?, update_at = ? WHERE id = %d`, repo.tableName, id)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		cake.Title,
		cake.Description,
		cake.Rating,
		cake.Image,
		cake.UpdateAt,
	)

	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}

func (repo *cakeRepositoryImpl) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = %d`, repo.tableName, id)
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
	)
	if err != nil {
		log.Println(err)
		return exception.ErrInternalServer
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		return exception.ErrNotFound
	}

	return nil
}
