package cake

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"privyID/helpers/response"
	"privyID/models"
)

type CakeHandler struct {
	Validate *validator.Validate
	UseCase  CakeUseCase
}

func NewCakeHandler(router *mux.Router, validate *validator.Validate, usecase CakeUseCase) {
	handler := &CakeHandler{
		Validate: validate,
		UseCase:  usecase,
	}

	router.HandleFunc("/cakes", handler.AddCakes).Methods(http.MethodPost)
	router.HandleFunc("/cakes/{id}", handler.DetailCakes).Methods(http.MethodGet)
	router.HandleFunc("/cakes", handler.ListCakes).Methods(http.MethodGet)
	router.HandleFunc("/cakes/{id}", handler.UpdateCake).Methods(http.MethodPatch)
	router.HandleFunc("/cakes/{id}", handler.DeleteCake).Methods(http.MethodDelete)
}

func (handler *CakeHandler) AddCakes(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput models.CheeseCake

	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		log.Println(err)
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.AddCakes(ctx, userInput)

	res.JSON(w)
}

func (handler *CakeHandler) DetailCakes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	res := handler.UseCase.DetailCakes(ctx, id)

	res.JSON(w)
}

func (handler *CakeHandler) ListCakes(w http.ResponseWriter, r *http.Request) {

	res := handler.UseCase.ListCakes()

	res.JSON(w)
}

func (handler *CakeHandler) UpdateCake(w http.ResponseWriter, r *http.Request) {
	var res response.Response
	var userInput models.CheeseCake

	ctx := r.Context()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		log.Println(err)
		res = response.Error(response.StatusUnprocessableEntity, err)
		res.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, userInput)
	if err != nil {
		res = response.Error(response.StatusBadRequest, err)
		res.JSON(w)
		return
	}

	res = handler.UseCase.UpdateCake(ctx, id, userInput)

	res.JSON(w)
}

func (handler *CakeHandler) DeleteCake(w http.ResponseWriter, r *http.Request) {
	var res response.Response

	ctx := r.Context()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	res = handler.UseCase.DeleteCake(ctx, id)

	res.JSON(w)
}
