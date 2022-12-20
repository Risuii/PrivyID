package cake_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"privyID/internal/cake"
	"privyID/internal/constant"
	"privyID/internal/mock"
	"privyID/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var currentTime = time.Date(2021, 12, 12, 0, 0, 0, 0, &time.Location{})

var kue = models.CheeseCake{
	ID:          1,
	Title:       "test-1",
	Description: "hanya untuk test",
	Rating:      1,
	Image:       "ini gambar test",
	CreatedAt:   currentTime,
	UpdateAt:    currentTime,
}

func TestCreate(t *testing.T) {
	t.Run("Test Create Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(kue.Title, kue.Description, kue.Rating, kue.Image, kue.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

		ID, err := repo.Create(ctx, kue)

		assert.Equal(t, int64(1), ID)
		assert.NoError(t, err)
	})

	t.Run("Test Create Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`INSERT INTO %s`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(kue.Title, kue.Description, kue.Rating, kue.Image, kue.CreatedAt).WillReturnResult(sqlmock.NewResult(0, 0))

		ID, err := repo.Create(ctx, kue)

		assert.Equal(t, int64(0), ID)
		assert.NoError(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Test Find By Id Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s WHERE id = ?`, constant.TableCake)
		rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "update_at"}).AddRow(kue.ID, kue.Title, kue.Description, kue.Rating, kue.Image, kue.CreatedAt, kue.UpdateAt)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(kue.ID).WillReturnRows(rows)

		kue, err := repo.FindByID(ctx, int64(kue.ID))

		assert.NotNil(t, kue)
		assert.NoError(t, err)
	})

	t.Run("Test Find By Id Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s WHERE id = ?`, constant.TableCake)
		rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "update_at"})

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectQuery().WithArgs(kue.ID).WillReturnRows(rows)

		kue, err := repo.FindByID(ctx, int64(kue.ID))
		assert.Empty(t, kue)
		assert.Error(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Test Find All Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s ORDER BY rating DESC, title ASC`, constant.TableCake)
		rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "update_at"}).AddRow(kue.ID, kue.Title, kue.Description, kue.Rating, kue.Image, kue.CreatedAt, kue.UpdateAt)

		mock.ExpectQuery(query).WillReturnRows(rows)

		kue, err := repo.FindAll()

		assert.NotNil(t, kue)
		assert.NoError(t, err)
	})

	t.Run("Test Find All Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`SELECT id, title, description, rating, image, created_at, update_at FROM %s ORDER BY rating DESC, title ASC`, constant.TableCake)
		rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "update_at"})

		mock.ExpectQuery(query).WillReturnRows(rows)

		kue, err := repo.FindAll()

		assert.Empty(t, kue)
		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Update Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(kue.Title, kue.Description, kue.Rating, kue.Image, kue.UpdateAt).WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.Update(ctx, int64(kue.ID), kue)

		assert.NoError(t, err)
	})

	t.Run("Test Update Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`UPDATE %s SET`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs(kue.Title, kue.Description, kue.Rating, kue.Image, kue.UpdateAt).WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Update(ctx, int64(kue.ID), kue)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Test Delete Success", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(ctx, int64(kue.ID))

		assert.NoError(t, err)
	})

	t.Run("Test Delete Error", func(t *testing.T) {
		db, mock := mock.NewMock()
		repo := cake.NewCakeRepository(db, constant.TableCake)

		defer db.Close()

		query := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, constant.TableCake)

		ctx := context.TODO()

		mock.ExpectPrepare(query).ExpectExec().WithArgs().WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.Delete(ctx, int64(kue.ID))

		assert.Error(t, err)
	})
}
