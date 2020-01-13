package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Royalsspirit/cisco-api/internal/api/models"
	"github.com/Royalsspirit/cisco-api/internal/api/util"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func (s *Server) list(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	characters, err := models.AllPeoples(s.DB)

	//Add business logic to characters items if needed

	if err != nil {
		util.ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (s *Server) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := json.NewDecoder(r.Body)

	var input models.People

	err := body.Decode(&input)
	/**
	* if body decoding failed stop execution
	**/
	if err != nil {
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	// Init validator
	validate, trans := s.Validator.Val, s.Validator.Trans

	validateErr := validate.Struct(input)

	if validateErr != nil {

		errors := util.BuildErrorsFromValidationErrors(validateErr.(validator.ValidationErrors), &trans)

		util.ResponseErrors(w, errors, http.StatusBadRequest)
		return
	}

	err = models.PeopleUpdateHandler(s.DB, input, vars["id"])

	if err != nil {
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)

	var input models.People

	err := body.Decode(&input)
	/**
	* if body decoding failed stop execution
	**/
	if err != nil {
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	// Init validator
	validate, trans := s.Validator.Val, s.Validator.Trans

	validateErr := validate.Struct(input)

	if validateErr != nil {

		errors := util.BuildErrorsFromValidationErrors(validateErr.(validator.ValidationErrors), &trans)

		util.ResponseErrors(w, errors, http.StatusBadRequest)
		return
	}

	err = models.InsertPeople(s.DB, input)

	if err != nil {
		fmt.Println("err", err)
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	err := models.DeletePeople(s.DB, vars["id"])

	if err != nil {
		fmt.Println("err", err)
		util.ResponseError(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
