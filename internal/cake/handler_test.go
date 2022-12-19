package cake

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"privyID/helpers/response"
	"privyID/internal/cake/mocks"
	"privyID/models"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCakeHandler_AddCakes(t *testing.T) {
	t.Run("add cakes success", func(t *testing.T) {
		newCake := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
		}

		newCakeRes := models.CheeseCake{
			ID:          newCake.ID,
			Title:       newCake.Title,
			Description: newCake.Description,
			Rating:      newCake.Rating,
			Image:       newCake.Image,
			CreatedAt:   time.Now(),
		}

		resp := response.Success(response.StatusCreated, newCakeRes)

		validate := validator.New()
		cakeUseCase := new(mocks.MockCake)
		cakeUseCase.On("AddCakes", mock.Anything, mock.AnythingOfType("models.CheeseCake")).Return(resp)

		newAddCakes, _ := json.Marshal(newCake)

		cakeHandler := CakeHandler{
			Validate: validate,
			UseCase:  cakeUseCase,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newAddCakes))
		recorder := httptest.NewRecorder()

		fmt.Println("INI FMT", r)

		handler := http.HandlerFunc(cakeHandler.AddCakes)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusCreated, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusCreated))
		assert.NotNil(t, rb.Data, "should not be nil")

		cakeUseCase.AssertExpectations(t)
	})

	t.Run("add cake failed", func(t *testing.T) {
		type invalidReq struct {
			Data string
		}

		newCakeReq := invalidReq{
			Data: "error",
		}

		validate := validator.New()
		cakeUseCaseTest := new(mocks.MockCake)

		newAddCakes, _ := json.Marshal(newCakeReq)

		cakeHandler := CakeHandler{
			Validate: validate,
			UseCase:  cakeUseCaseTest,
		}

		r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newAddCakes))
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(cakeHandler.AddCakes)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}
		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusBadRequest, rb.Status, "should be bad request")
		assert.Nil(t, rb.Data, "should be nil")

		cakeUseCaseTest.AssertExpectations(t)
	})
}

func TestCakeHandler_DetailCakes(t *testing.T) {
	t.Run("Get DetailCakes Success", func(t *testing.T) {
		newCake := models.CheeseCake{
			ID:          1,
			Title:       "test-1",
			Description: "hanya untuk test",
			Rating:      1,
			Image:       "ini gambar test",
			CreatedAt:   time.Time{},
		}

		resp := response.Success(response.StatusOK, newCake)

		cakeUseCaseTest := new(mocks.MockCake)
		cakeUseCaseTest.On("DetailCakes", mock.Anything, mock.Anything).Return(resp)

		CakeHandler := CakeHandler{
			UseCase: cakeUseCaseTest,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(CakeHandler.DetailCakes)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}

		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		// rb.Data = newCake
		// rb.Status = response.StatusOK

		assert.Equal(t, response.StatusOK, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusOK))
		assert.NotNil(t, rb.Data, "should not be nil")

		cakeUseCaseTest.AssertExpectations(t)
	})

	t.Run("Get DetailCakes Failed", func(t *testing.T) {

		resp := response.Error(response.StatusInternalServerError, assert.AnError)

		cakeUseCaseTest := new(mocks.MockCake)
		cakeUseCaseTest.On("DetailCakes", mock.Anything, mock.Anything).Return(resp)

		CakeHandler := CakeHandler{
			UseCase: cakeUseCaseTest,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(CakeHandler.DetailCakes)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}

		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusInternalServerError, rb.Status, "should be internal server error")
		assert.Nil(t, rb.Data, "should be nil")

		cakeUseCaseTest.AssertExpectations(t)
	})
}

func TestCakeHandler_ListCakes(t *testing.T) {
	t.Run("Get ListCakes Success", func(t *testing.T) {
		newCake := []models.CheeseCake{
			{
				ID:          1,
				Title:       "test-1",
				Description: "hanya untuk test",
				Rating:      1,
				Image:       "ini gambar test",
				CreatedAt:   time.Time{},
			},
		}

		resp := response.Success(response.StatusOK, newCake)

		cakeUseCaseTest := new(mocks.MockCake)
		cakeUseCaseTest.On("ListCakes").Return(resp)

		CakeHandler := CakeHandler{
			UseCase: cakeUseCaseTest,
		}

		r := httptest.NewRequest(http.MethodGet, "/just/for/testing", nil)
		recorder := httptest.NewRecorder()

		handler := http.HandlerFunc(CakeHandler.ListCakes)
		handler.ServeHTTP(recorder, r)

		rb := response.ResponseImpl{}

		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, response.StatusOK, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusOK))
		assert.NotNil(t, rb.Data, "should not be nil")

		cakeUseCaseTest.AssertExpectations(t)
	})

	t.Run("Get ListCakes Failed", func(t *testing.T) {

	})
}

// func TestCakeHandler_UpdateCake(t *testing.T) {
// 	t.Run("Update Cake Success", func(t *testing.T) {
// 		newCake := models.CheeseCake{
// 			ID:          0,
// 			Title:       "test-1",
// 			Description: "hanya untuk test",
// 			Rating:      1,
// 			Image:       "ini gambar test",
// 		}

// 		resp := response.Success(response.StatusOK, newCake)

// 		validate := validator.New()
// 		cakeUseCaseTest := new(mocks.MockCake)
// 		cakeUseCaseTest.On("UpdateCake", mock.Anything, int64(0), mock.Anything).Return(resp)

// 		newCakeRes, err := json.Marshal(newCake)
// 		if err != nil {
// 			log.Println(err)
// 			t.Error(err)
// 			return
// 		}

// 		cakeHandlerTest := CakeHandler{
// 			Validate: validate,
// 			UseCase:  cakeUseCaseTest,
// 		}

// 		r := httptest.NewRequest(http.MethodPatch, "/just/for/testing", bytes.NewReader(newCakeRes))
// 		recorder := httptest.NewRecorder()

// 		handler := http.HandlerFunc(cakeHandlerTest.UpdateCake)
// 		handler.ServeHTTP(recorder, r)

// 		rb := response.ResponseImpl{}
// 		if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
// 			log.Println(err)
// 			t.Error(err)
// 			return
// 		}

// 		assert.Equal(t, response.StatusOK, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusOK))
// 		assert.NotNil(t, rb.Data, "should be nil")

// 		cakeUseCaseTest.AssertExpectations(t)
// 	})
// }
