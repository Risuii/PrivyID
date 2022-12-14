package cake

import (
	"context"
	"log"
	"time"

	"privyID/helpers/exception"
	"privyID/helpers/response"
	"privyID/models"
)

type (
	CakeUseCase interface {
		AddCakes(ctx context.Context, params models.CheeseCake) response.Response
		DetailCakes(ctx context.Context, id int64) response.Response
		ListCakes() response.Response
		UpdateCake(ctx context.Context, id int64, params models.CheeseCake) response.Response
		DeleteCake(ctx context.Context, id int64) response.Response
	}

	cakeUseCaseImpl struct {
		repository CakeRepository
	}
)

func NewCakeUseCase(repo CakeRepository) CakeUseCase {
	return &cakeUseCaseImpl{
		repository: repo,
	}
}

func (cu *cakeUseCaseImpl) AddCakes(ctx context.Context, params models.CheeseCake) response.Response {
	cakes := models.CheeseCake{
		ID:          params.ID,
		Title:       params.Title,
		Description: params.Description,
		Rating:      params.Rating,
		Image:       params.Image,
		CreatedAt:   time.Now(),
	}

	cakeID, err := cu.repository.Create(ctx, cakes)
	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	cakes.ID = int(cakeID)

	return response.Success(response.StatusCreated, cakes)
}

func (cu *cakeUseCaseImpl) DetailCakes(ctx context.Context, id int64) response.Response {
	cakes, err := cu.repository.FindByID(ctx, id)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, cakes)
}

func (cu *cakeUseCaseImpl) ListCakes() response.Response {

	cakeData, err := cu.repository.FindAll()
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, cakeData)
}

func (cu *cakeUseCaseImpl) UpdateCake(ctx context.Context, id int64, params models.CheeseCake) response.Response {
	var cakes models.CheeseCake

	cakeId, err := cu.repository.FindByID(ctx, id)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}
	if err != nil {
		log.Println(err)
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	cakes.Title = params.Title
	cakes.Description = params.Description
	cakes.Rating = params.Rating
	cakes.Image = params.Image
	cakes.UpdateAt = time.Now()

	err = cu.repository.Update(ctx, int64(cakeId.ID), cakes)

	if err != nil {
		log.Println(err)
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	cakes.ID = cakeId.ID

	return response.Success(response.StatusOK, cakes)
}

func (cu *cakeUseCaseImpl) DeleteCake(ctx context.Context, id int64) response.Response {

	err := cu.repository.Delete(ctx, id)
	if err == exception.ErrNotFound {
		return response.Error(response.StatusNotFound, exception.ErrNotFound)
	}

	if err != nil {
		return response.Error(response.StatusInternalServerError, exception.ErrInternalServer)
	}

	msg := "Success Delete Data"
	return response.Success(response.StatusOK, msg)
}
