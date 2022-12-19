package cake_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"privyID/helpers/exception"
	"privyID/internal/cake"
	"privyID/internal/cake/mocks"
	"privyID/models"
)

func TestUseCaseAddCake(t *testing.T) {
	t.Run("Add Cake Success", func(t *testing.T) {
		addCakeRepository := new(mocks.CakeRepository)

		addCakeRepository.On("Create", mock.Anything, mock.AnythingOfType("models.CheeseCake")).Return(int64(1), nil)

		cakeUseCaseTest := cake.NewCakeUseCase(
			addCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCaseTest.AddCakes(ctx, params)

		assert.NoError(t, resp.Err())

		addCakeRepository.AssertExpectations(t)
	})

	t.Run("Add Cake Error", func(t *testing.T) {
		addCakeRepository := new(mocks.CakeRepository)

		addCakeRepository.On("Create", mock.Anything, mock.AnythingOfType("models.CheeseCake")).Return(int64(0), exception.ErrInternalServer)

		cakeUseCaseTest := cake.NewCakeUseCase(
			addCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCaseTest.AddCakes(ctx, params)

		assert.Error(t, resp.Err())

		addCakeRepository.AssertExpectations(t)
	})
}

func TestUseCaseDetailCakes(t *testing.T) {
	t.Run("Get Detail Cake Success", func(t *testing.T) {
		detailCakeRepository := new(mocks.CakeRepository)

		detailCakeRepository.On("FindByID", mock.Anything, mock.Anything).Return(models.CheeseCake{}, nil)

		cakeUseCaseTest := cake.NewCakeUseCase(
			detailCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCaseTest.DetailCakes(ctx, int64(params.ID))

		assert.NoError(t, resp.Err())

		detailCakeRepository.AssertExpectations(t)
	})

	t.Run("Get Detail Cake Error", func(t *testing.T) {
		detailCakeRepository := new(mocks.CakeRepository)

		detailCakeRepository.On("FindByID", mock.Anything, mock.Anything).Return(models.CheeseCake{}, exception.ErrNotFound)

		cakeUseCaseTest := cake.NewCakeUseCase(
			detailCakeRepository,
		)

		ctx := context.TODO()

		resp := cakeUseCaseTest.DetailCakes(ctx, 0)

		assert.Error(t, resp.Err())

		detailCakeRepository.AssertExpectations(t)
	})

	t.Run("Get Detail Cake Internal Server Error", func(t *testing.T) {
		detailCakeRepository := new(mocks.CakeRepository)

		detailCakeRepository.On("FindByID", mock.Anything, mock.Anything).Return(models.CheeseCake{}, exception.ErrInternalServer)

		cakeUseCaseTest := cake.NewCakeUseCase(
			detailCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCaseTest.DetailCakes(ctx, int64(params.ID))

		assert.Error(t, resp.Err())

		detailCakeRepository.AssertExpectations(t)

	})
}

func TestUseCaseListCake(t *testing.T) {
	t.Run("Get List Cake Success", func(t *testing.T) {
		listCakeRepository := new(mocks.CakeRepository)

		listCakeRepository.On("FindAll").Return([]models.CheeseCake{}, nil)

		cakeUseCaseTest := cake.NewCakeUseCase(
			listCakeRepository,
		)

		resp := cakeUseCaseTest.ListCakes()

		assert.NoError(t, resp.Err())
		listCakeRepository.AssertExpectations(t)
	})

	t.Run("Get List Cake Error", func(t *testing.T) {
		listCakeRepository := new(mocks.CakeRepository)

		listCakeRepository.On("FindAll").Return([]models.CheeseCake{}, exception.ErrNotFound)

		cakeUseCaseTest := cake.NewCakeUseCase(
			listCakeRepository,
		)

		resp := cakeUseCaseTest.ListCakes()

		assert.Error(t, resp.Err())
		listCakeRepository.AssertExpectations(t)
	})

	t.Run("Get List Cake Internal Error", func(t *testing.T) {
		listCakeRepository := new(mocks.CakeRepository)

		listCakeRepository.On("FindAll").Return([]models.CheeseCake{}, exception.ErrInternalServer)

		cakeUseCaseTest := cake.NewCakeUseCase(
			listCakeRepository,
		)

		resp := cakeUseCaseTest.ListCakes()

		assert.Error(t, resp.Err())
		listCakeRepository.AssertExpectations(t)
	})
}

func TestUseCaseUpdateCake(t *testing.T) {

	t.Run("Update Cake Success", func(t *testing.T) {
		updateCakeRepository := &mocks.CakeRepository{}
		cakeUseCaseTest := cake.NewCakeUseCase(
			updateCakeRepository,
		)

		// Menyiapkan comportemen mock
		updateCakeRepository.On("FindByID", mock.Anything, int64(1)).Return(models.CheeseCake{ID: 1}, nil)
		updateCakeRepository.On("Update", mock.Anything, int64(1), mock.Anything).Return(nil)

		ctx := context.TODO()
		params := models.CheeseCake{
			Title:       "Chocolate Cake",
			Description: "Delicious chocolate cake",
			Rating:      5,
			Image:       "chocolate_cake.jpg",
		}

		// Memanggil fungsi UpdateCake dengan input yang sesuai
		res := cakeUseCaseTest.UpdateCake(ctx, 1, params)

		assert.NoError(t, res.Err())
		updateCakeRepository.AssertExpectations(t)
	})

	t.Run("Update Cake Error", func(t *testing.T) {
		updateCakeRepository := new(mocks.CakeRepository)

		updateCakeRepository.On("FindByID", mock.Anything, mock.Anything).Return(models.CheeseCake{}, exception.ErrNotFound)

		cakeUseCase := cake.NewCakeUseCase(
			updateCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCase.UpdateCake(ctx, int64(params.ID), params)

		assert.Error(t, resp.Err())

		updateCakeRepository.AssertExpectations(t)
	})

	t.Run("Update Cake Internal Server Error", func(t *testing.T) {
		updateCakeRepository := new(mocks.CakeRepository)

		updateCakeRepository.On("FindByID", mock.Anything, mock.Anything).Return(models.CheeseCake{}, exception.ErrInternalServer)

		cakeUseCase := cake.NewCakeUseCase(
			updateCakeRepository,
		)

		ctx := context.TODO()
		params := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		resp := cakeUseCase.UpdateCake(ctx, int64(params.ID), params)

		assert.Error(t, resp.Err())

		updateCakeRepository.AssertExpectations(t)
	})
}

func TestDeleteCake(t *testing.T) {
	t.Run("Delete Cake Success", func(t *testing.T) {
		deleteCakeRepository := &mocks.CakeRepository{}
		cakeUseCaseTest := cake.NewCakeUseCase(
			deleteCakeRepository,
		)

		deleteCakeRepository.On("Delete", mock.Anything, int64(1)).Return(nil)

		ctx := context.TODO()

		res := cakeUseCaseTest.DeleteCake(ctx, 1)
		assert.NoError(t, res.Err())
		deleteCakeRepository.AssertExpectations(t)
	})

	t.Run("Delete Cake Error Not Found", func(t *testing.T) {
		deleteCakeRepository := &mocks.CakeRepository{}
		cakeUseCaseTest := cake.NewCakeUseCase(
			deleteCakeRepository,
		)

		deleteCakeRepository.On("Delete", mock.Anything, int64(9999)).Return(exception.ErrNotFound)

		ctx := context.TODO()

		res := cakeUseCaseTest.DeleteCake(ctx, 9999)

		assert.Error(t, res.Err())
		deleteCakeRepository.AssertExpectations(t)
	})

	t.Run("Delete Cake Error internal server", func(t *testing.T) {
		deleteCakeRepository := &mocks.CakeRepository{}
		cakeUseCaseTest := cake.NewCakeUseCase(
			deleteCakeRepository,
		)

		deleteCakeRepository.On("Delete", mock.Anything, int64(1)).Return(exception.ErrInternalServer)

		ctx := context.TODO()

		res := cakeUseCaseTest.DeleteCake(ctx, 1)

		assert.Error(t, res.Err())
		deleteCakeRepository.AssertExpectations(t)
	})
}
