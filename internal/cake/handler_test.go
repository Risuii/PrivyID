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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCakeHandler_AddCakes(t *testing.T) {
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

	cakeUseCase := new(mocks.MockCake)
	cakeUseCase.On("AddCakes", mock.Anything, mock.AnythingOfType("models.CheeseCake")).Return(resp)

	newAddCakes, _ := json.Marshal(newCake)

	cakeHandler := CakeHandler{
		UseCase: cakeUseCase,
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

	assert.Equal(t, response.StatusCreated, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusCreated))
	assert.NotNil(t, rb.Data, "should not be nil")

	cakeUseCase.AssertExpectations(t)
}

func TestCakeHandler_AddCakes_Failed(t *testing.T) {
	type invalidReq struct {
		Data string
	}

	newCakeReq := invalidReq{
		Data: "error",
	}

	cakeUseCaseTest := new(mocks.MockCake)

	newAddCakes, _ := json.Marshal(newCakeReq)

	cakeHandler := CakeHandler{
		UseCase: cakeUseCaseTest,
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

	assert.Equal(t, response.StatusUnprocessableEntity, rb.Status, "should be bad request")
	assert.Nil(t, rb.Data, "should be nil")

	cakeUseCaseTest.AssertExpectations(t)
}
